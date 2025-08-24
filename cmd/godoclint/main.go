package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/godoc-lint/godoc-lint/pkg/analysis"
	"github.com/godoc-lint/godoc-lint/pkg/check"
	"github.com/godoc-lint/godoc-lint/pkg/config"
	"github.com/godoc-lint/godoc-lint/pkg/inspect"
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

	reg := check.NewPopulatedRegistry()
	cb := config.NewConfigBuilder(baseDir, reg.GetCoveredRules())
	ocb := config.NewOnceConfigBuilder(cb)
	inspector := inspect.NewInspector(ocb, exitFunc)
	analyzer := analysis.NewAnalyzer(baseDir, ocb, reg, inspector, exitFunc)

	configOverride := model.NewConfigOverride()
	ocb.SetOverride(configOverride)

	analyzer.GetAnalyzer().Flags.Func("config", "path to config file", func(s string) error {
		if configOverride.ConfigFilePath != nil {
			return errors.New("config file is set multiple times")
		}
		if strings.TrimSpace(s) == "" {
			return errors.New("empty path")
		}
		configOverride.ConfigFilePath = &s
		return nil
	})

	walkNonEmptyCSV := func(f func(string)) func(string) error {
		return func(value string) error {
			values := strings.Split(strings.TrimSpace(value), ",")
			for _, v := range values {
				if strings.TrimSpace(v) == "" {
					return errors.New("empty element")
				}
				f(v)
			}
			return nil
		}
	}

	walkNonEmpty := func(f func(string)) func(string) error {
		return func(value string) error {
			if strings.TrimSpace(value) == "" {
				return errors.New("empty value")
			}
			f(value)
			return nil
		}
	}

	analyzer.GetAnalyzer().Flags.Func("include", "regexp path (Unix style) to include (can be used multiple times)", walkNonEmpty(func(s string) {
		configOverride.Include = append(configOverride.Include, s)
	}))
	analyzer.GetAnalyzer().Flags.Func("exclude", "regexp path (Unix style) to exclude (can be used multiple times)", walkNonEmpty(func(s string) {
		configOverride.Exclude = append(configOverride.Exclude, s)
	}))
	analyzer.GetAnalyzer().Flags.Func("enable", "comma-separated rule names to enable", walkNonEmptyCSV(func(s string) {
		configOverride.Enable = append(configOverride.Enable, s)
	}))
	analyzer.GetAnalyzer().Flags.Func("disable", "comma-separated rule names to disable", walkNonEmptyCSV(func(s string) {
		configOverride.Disable = append(configOverride.Disable, s)
	}))

	analyzer.GetAnalyzer().Flags.BoolFunc("V", "print version and exit", func(s string) error {
		fmt.Println(version.Current)
		os.Exit(0)
		return nil
	})

	singlechecker.Main(analyzer.GetAnalyzer())
}
