package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/godoc-lint/godoc-lint/pkg/model"
	"github.com/godoc-lint/godoc-lint/pkg/util"
)

// ConfigBuilder implements a configuration builder.
type ConfigBuilder struct {
	baseDir      string
	coveredRules model.RuleSet
	override     *model.ConfigOverride
}

// NewConfigBuilder crates a new instance of the corresponding struct.
func NewConfigBuilder(baseDir string, coveredRules model.RuleSet) *ConfigBuilder {
	return &ConfigBuilder{
		baseDir:      baseDir,
		coveredRules: coveredRules,
	}
}

// GetConfig implements the corresponding interface method.
func (cb *ConfigBuilder) GetConfig(cwd string) (model.Config, error) {
	return cb.build(cwd)
}

func (cb *ConfigBuilder) resolvePlainConfig(cwd string) (*PlainConfig, *PlainConfig, string, error) {
	def, err := FromYAML(defaultConfigYAML)
	if err != nil {
		// This should never happen.
		panic("cannot parse default config")
	}

	if !util.IsPathUnderBaseDir(cb.baseDir, cwd) {
		if pcfg, err := cb.resolvePlainConfigAtBaseDir(); err != nil {
			return nil, nil, "", err
		} else if pcfg != nil {
			return pcfg, def, cb.baseDir, nil
		}
		return def, def, cb.baseDir, nil
	}

	path := cwd
	for {
		rel, err := filepath.Rel(cb.baseDir, path)
		if err != nil {
			return nil, nil, "", err
		}

		if rel == "." {
			if pcfg, err := cb.resolvePlainConfigAtBaseDir(); err != nil {
				return nil, nil, "", err
			} else if pcfg != nil {
				return pcfg, def, cb.baseDir, nil
			}
			return def, def, cb.baseDir, nil
		}

		if pcfg, err := findConventionalConfigFile(path); err != nil {
			return nil, nil, "", err
		} else if pcfg != nil {
			return pcfg, def, path, nil
		}

		path = filepath.Dir(path)
	}
}

func (cb *ConfigBuilder) resolvePlainConfigAtBaseDir() (*PlainConfig, error) {
	if cb.override == nil || cb.override.ConfigFilePath == nil {
		return findConventionalConfigFile(cb.baseDir)
	}
	pcfg, err := FromYAMLFile(*cb.override.ConfigFilePath)
	if err != nil {
		return nil, err
	}
	return pcfg, nil
}

func findConventionalConfigFile(dir string) (*PlainConfig, error) {
	for _, dcf := range defaultConfigFiles {
		path := filepath.Join(dir, dcf)
		if fi, err := os.Stat(path); err != nil || fi.IsDir() {
			continue
		}
		pcfg, err := FromYAMLFile(path)
		if err != nil {
			return nil, err
		}
		return pcfg, nil
	}
	return nil, nil
}

// build creates the configuration struct.
//
// It checks a sequence of sources:
//  1. Custom config file path
//  2. Default configuration files (e.g., `.godoc-lint.yaml`)
//
// If none was available, the default configuration will be returned.
//
// The method also does the following:
//   - Applies override flags (e.g., enable, or disable).
//   - Validates the final configuration.
func (cb *ConfigBuilder) build(cwd string) (*config, error) {
	pcfg, def, configCWD, err := cb.resolvePlainConfig(cwd)
	if err != nil {
		return nil, err
	}

	toValidRuleSet := func(s []string) (*model.RuleSet, []string) {
		if s == nil {
			return nil, nil
		}
		invalids := make([]string, 0, len(s))
		rules := make([]string, 0, len(s))
		for _, v := range s {
			if !cb.coveredRules.Has(v) {
				invalids = append(invalids, v)
				continue
			}
			rules = append(rules, v)
		}
		set := model.RuleSet{}.Add(rules...)
		return &set, invalids
	}

	toValidRegexpSlice := func(s []string) ([]*regexp.Regexp, []string) {
		if s == nil {
			return nil, nil
		}
		var invalids []string
		var regexps []*regexp.Regexp
		for _, v := range s {
			re, err := regexp.Compile(v)
			if err != nil {
				invalids = append(invalids, v)
				continue
			}
			regexps = append(regexps, re)
		}
		return regexps, invalids
	}

	var errs error

	resolvedEnable := pcfg.Enable
	if cb.override != nil && cb.override.Enable != nil {
		resolvedEnable = cb.override.Enable
	}
	if resolvedEnable == nil {
		resolvedEnable = def.Enable
	}
	resolvedEnabledRuleSet, invalids := toValidRuleSet(resolvedEnable)
	if len(invalids) > 0 {
		errs = errors.Join(errs, fmt.Errorf("invalid rule(s) name to enable: %q", invalids))
	}

	resolvedDisable := pcfg.Disable
	if cb.override != nil && cb.override.Disable != nil {
		resolvedDisable = cb.override.Disable
	}
	if resolvedDisable == nil {
		resolvedDisable = def.Disable
	}
	resolvedDisabledRuleSet, invalids := toValidRuleSet(resolvedDisable)
	if len(invalids) > 0 {
		errs = errors.Join(errs, fmt.Errorf("invalid rule(s) to disable: %q", invalids))
	}

	resolvedInclude := pcfg.Include
	if cb.override != nil && cb.override.Include != nil {
		resolvedInclude = cb.override.Include
	}
	if resolvedInclude == nil {
		resolvedInclude = def.Include
	}
	resolvedIncludeAsRegexp, invalids := toValidRegexpSlice(resolvedInclude)
	if len(invalids) > 0 {
		errs = errors.Join(errs, fmt.Errorf("invalid path pattern(s) to include: %q", invalids))
	}

	resolvedExclude := pcfg.Exclude
	if cb.override != nil && cb.override.Exclude != nil {
		resolvedExclude = cb.override.Exclude
	}
	if resolvedExclude == nil {
		resolvedExclude = def.Exclude
	}
	resolvedExcludeAsRegexp, invalids := toValidRegexpSlice(resolvedExclude)
	if len(invalids) > 0 {
		errs = errors.Join(errs, fmt.Errorf("invalid path pattern(s) to exclude: %q", invalids))
	}

	if errs != nil {
		return nil, errs
	}

	resolvedOptions := &model.RuleOptions{}
	transferOptions(resolvedOptions, def.Options) // def.Options is never nil
	if pcfg.Options != nil {
		transferOptions(resolvedOptions, pcfg.Options)
	}

	return &config{
		cwd: configCWD,

		enabledRules:    resolvedEnabledRuleSet,
		disabledRules:   resolvedDisabledRuleSet,
		includeAsRegexp: resolvedIncludeAsRegexp,
		excludeAsRegexp: resolvedExcludeAsRegexp,
		options:         resolvedOptions,
	}, nil
}

// SetOverride implements the corresponding interface method.
func (cb *ConfigBuilder) SetOverride(override *model.ConfigOverride) {
	cb.override = override
}
