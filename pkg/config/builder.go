package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/godoc-lint/godoc-lint/pkg/model"
)

// ConfigBuilder implements a configuration builder.
type ConfigBuilder struct {
	cwd          string
	coveredRules model.RuleSet
	override     *model.ConfigOverride
}

// NewConfigBuilder crates a new instance of the corresponding struct.
func NewConfigBuilder(cwd string, coveredRules model.RuleSet) *ConfigBuilder {
	return &ConfigBuilder{
		cwd:          cwd,
		coveredRules: coveredRules,
	}
}

// GetConfig implements the corresponding interface method.
func (cb *ConfigBuilder) GetConfig(path string) (model.Config, error) {
	return cb.build(path)
}

func (cb *ConfigBuilder) resolvePlainConfig(cwd string) (*PlainConfig, *PlainConfig, error) {
	def, err := FromYAML(defaultConfigYAML)
	if err != nil {
		// This should never happen.
		panic("cannot parse default config")
	}

	rel, err := filepath.Rel(cb.cwd, cwd)
	if err != nil {
		return nil, nil, fmt.Errorf("cannot find relative path to base dir: %w", err)
	}

	// The config file path override should only be applied to the root config
	// or if the given path is not under the original CWD.
	if rel == "." && cb.override != nil && cb.override.ConfigFilePath != nil {
		pcfg, err := FromYAMLFile(*cb.override.ConfigFilePath)
		if err != nil {
			return nil, nil, fmt.Errorf("cannot read config file (%q): %w", *cb.override.ConfigFilePath, err)
		}
		return pcfg, def, nil
	}

	if rel == ".." || strings.HasPrefix(filepath.ToSlash(rel), "../") {
		// The given CWD is outside the original CWD. So, we apply the default.
		return def, def, nil
	}

	path := cwd
	for {
		for _, dcf := range defaultConfigFiles {
			p := filepath.Join(path, dcf)
			if fi, err := os.Stat(p); err != nil || fi.IsDir() {
				continue
			}
			pcfg, err := FromYAMLFile(p)
			if err != nil {
				return nil, nil, fmt.Errorf("malformed configuration file (at %q): %w", p, err)
			}
			return pcfg, def, nil
		}

		if rel, err := filepath.Rel(cb.cwd, path); err != nil {
			return nil, nil, fmt.Errorf("cannot find relative path to base dir: %w", err)
		} else if rel[0] == '.' {
			break
		}
		path = filepath.Dir(path)
	}
	return def, def, nil
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
	pcfg, def, err := cb.resolvePlainConfig(cwd)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
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
		errs = errors.Join(errs, fmt.Errorf("config error: invalid rule(s) name to enable: %q", invalids))
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
		errs = errors.Join(errs, fmt.Errorf("config error: invalid rule(s) to disable: %q", invalids))
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
		errs = errors.Join(errs, fmt.Errorf("config error: invalid path pattern(s) to include: %q", invalids))
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
		errs = errors.Join(errs, fmt.Errorf("config error: invalid path pattern(s) to exclude: %q", invalids))
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
