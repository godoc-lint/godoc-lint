package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sync"

	"github.com/godoc-lint/godoc-lint/pkg/model"
)

// ConfigBuilder implements a configuration builder.
type ConfigBuilder struct {
	cwd          string
	coveredRules model.RuleSet
	exitFunc     func(int)

	override *model.ConfigOverride

	mu          sync.RWMutex
	built       bool
	builtConfig model.Config
	builtErr    error
}

// NewConfigBuilder crates a new instance of the corresponding struct.
func NewConfigBuilder(cwd string, coveredRules model.RuleSet, exitFunc func(int)) *ConfigBuilder {
	return &ConfigBuilder{
		cwd:          cwd,
		coveredRules: coveredRules,
		exitFunc:     exitFunc,
	}
}

// GetConfig implements the corresponding interface method.
func (cb *ConfigBuilder) GetConfig() (model.Config, error) {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	if !cb.built {
		cb.build()
	}
	return cb.builtConfig, cb.builtErr
}

// MustGetConfig implements the corresponding interface method.
func (cb *ConfigBuilder) MustGetConfig() model.Config {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	if !cb.built {
		cb.build()
	}

	if cb.builtErr != nil {
		fmt.Fprintln(os.Stderr, cb.builtErr.Error())
		cb.exit(2)
	}
	return cb.builtConfig
}

func (cb *ConfigBuilder) resolvePlainConfig() (*PlainConfig, *PlainConfig, error) {
	def, err := FromYAML(defaultConfigYAML)
	if err != nil {
		// This should never happen.
		panic("cannot parse default config")
	}

	if cb.override != nil && cb.override.ConfigFilePath != nil {
		pcfg, err := FromYAMLFile(*cb.override.ConfigFilePath)
		if err != nil {
			return nil, nil, fmt.Errorf("cannot read config file (%q): %w", *cb.override.ConfigFilePath, err)
		}
		return pcfg, def, nil
	}

	for _, dcf := range defaultConfigFiles {
		path := dcf
		if cb.cwd != "" {
			path = filepath.Join(cb.cwd, dcf)
		}

		fi, err := os.Stat(path)
		if err != nil || fi.IsDir() {
			continue
		}

		pcfg, err := FromYAMLFile(path)
		if err != nil {
			return nil, nil, fmt.Errorf("malformed configuration file (%q): %w", path, err)
		}
		return pcfg, def, nil
	}
	return def, def, nil
}

// build creates the configuration struct.
//
// If the provided inputs are not valid it will terminate the process.
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
//
// Caller should use a mutex to avoid concurrent calls.
func (cb *ConfigBuilder) build() {
	defer func() {
		cb.built = true
	}()

	pcfg, def, err := cb.resolvePlainConfig()
	if err != nil {
		cb.builtConfig = nil
		cb.builtErr = fmt.Errorf("config error: %w", err)
		return
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
		cb.builtConfig = nil
		cb.builtErr = errs
		return
	}

	built := &config{
		enabledRules:    resolvedEnabledRuleSet,
		disabledRules:   resolvedDisabledRuleSet,
		includeAsRegexp: resolvedIncludeAsRegexp,
		excludeAsRegexp: resolvedExcludeAsRegexp,
		options:         pcfg.extractRuleOptions(),
	}

	cb.builtConfig = built
	cb.builtErr = nil
}

func (cb *ConfigBuilder) exit(code int) {
	if cb.exitFunc != nil {
		cb.exitFunc(code)
		panic(fmt.Sprintf("exit code %d", code))
	}
	os.Exit(code)
}

// SetOverride implements the corresponding interface method.
func (cb *ConfigBuilder) SetOverride(override *model.ConfigOverride) {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	if cb.built {
		panic("config builder already sealed")
	}
	cb.override = override
}
