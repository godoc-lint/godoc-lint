// Godoc-Lint command package.
package main

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strings"

	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/godoc-lint/godoc-lint/pkg/compose"
	"github.com/godoc-lint/godoc-lint/pkg/model"
	"github.com/godoc-lint/godoc-lint/pkg/version"
)

func main() {
	exitFunc := func(code int, err error) {
		fmt.Fprintf(os.Stderr, "%v\n", err.Error())
		os.Exit(code)
	}

	baseDir, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to get current working directory: %v\n", err)
		os.Exit(1)
	}

	composition := compose.Compose(compose.CompositionConfig{
		BaseDir:  baseDir,
		ExitFunc: exitFunc,
	})

	configOverride := model.NewConfigOverride()
	composition.ConfigBuilder.SetOverride(configOverride)

	composition.Analyzer.GetAnalyzer().Flags.Func("config", "path to config file", func(s string) error {
		if configOverride.ConfigFilePath != nil {
			return errors.New("config file is set multiple times")
		}
		if strings.TrimSpace(s) == "" {
			return errors.New("empty path")
		}
		configOverride.ConfigFilePath = &s
		return nil
	})

	composition.Analyzer.GetAnalyzer().Flags.Func("default", "default set of rules to enable", func(s string) error {
		if configOverride.Default != nil {
			return errors.New("default set is set multiple times")
		}
		if !slices.Contains(model.DefaultSetValues, model.DefaultSet(s)) {
			return fmt.Errorf("unknown default set %q, must be one of %q", s, model.DefaultSetValues)
		}
		v := model.DefaultSet(s)
		configOverride.Default = &v
		return nil
	})

	walkNonEmptyCSV := func(f func(string) error) func(string) error {
		return func(value string) error {
			for v := range strings.SplitSeq(strings.TrimSpace(value), ",") {
				if strings.TrimSpace(v) == "" {
					return errors.New("empty element")
				}
				if err := f(v); err != nil {
					return err
				}
			}
			return nil
		}
	}

	walkNonEmpty := func(f func(string) error) func(string) error {
		return func(value string) error {
			if strings.TrimSpace(value) == "" {
				return errors.New("empty value")
			}
			if err := f(value); err != nil {
				return err
			}
			return nil
		}
	}

	composition.Analyzer.GetAnalyzer().Flags.Func("include", "regexp path (Unix style) to include (can be used multiple times)", walkNonEmpty(func(s string) error {
		re, err := regexp.Compile(s)
		if err != nil {
			return fmt.Errorf("invalid inclusion regexp pattern %q: %w", s, err)
		}
		configOverride.Include = append(configOverride.Include, re)
		return nil
	}))

	composition.Analyzer.GetAnalyzer().Flags.Func("exclude", "regexp path (Unix style) to exclude (can be used multiple times)", walkNonEmpty(func(s string) error {
		re, err := regexp.Compile(s)
		if err != nil {
			return fmt.Errorf("invalid exclusion regexp pattern %q: %w", s, err)
		}
		configOverride.Exclude = append(configOverride.Exclude, re)
		return nil
	}))

	composition.Analyzer.GetAnalyzer().Flags.Func("enable", "comma-separated rule names to enable", walkNonEmptyCSV(func(s string) error {
		if !model.AllRules.Has(model.Rule(s)) {
			return fmt.Errorf("unknown rule name to enable %q", s)
		}
		if configOverride.Disable != nil && configOverride.Disable.Has(model.Rule(s)) {
			return fmt.Errorf("cannot enable and disable rule at the same time %q", s)
		}
		if configOverride.Enable == nil {
			configOverride.Enable = &model.RuleSet{}
		}
		updated := configOverride.Enable.Add(model.Rule(s))
		configOverride.Enable = &updated
		return nil
	}))

	composition.Analyzer.GetAnalyzer().Flags.Func("disable", "comma-separated rule names to disable", walkNonEmptyCSV(func(s string) error {
		if !model.AllRules.Has(model.Rule(s)) {
			return fmt.Errorf("unknown rule name to disable %q", s)
		}
		if configOverride.Enable != nil && configOverride.Enable.Has(model.Rule(s)) {
			return fmt.Errorf("cannot enable and disable rule at the same time %q", s)
		}
		if configOverride.Disable == nil {
			configOverride.Disable = &model.RuleSet{}
		}
		updated := configOverride.Disable.Add(model.Rule(s))
		configOverride.Disable = &updated
		return nil
	}))

	composition.Analyzer.GetAnalyzer().Flags.BoolFunc("V", "print version and exit", func(s string) error {
		fmt.Println(version.Current)
		os.Exit(0)
		return nil
	})

	singlechecker.Main(composition.Analyzer.GetAnalyzer())
}
