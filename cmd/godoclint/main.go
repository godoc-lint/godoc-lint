package main

import (
	"errors"
	"fmt"
	"os"
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

	composition.Analyzer.GetAnalyzer().Flags.Func("include", "regexp path (Unix style) to include (can be used multiple times)", walkNonEmpty(func(s string) {
		configOverride.Include = append(configOverride.Include, s)
	}))
	composition.Analyzer.GetAnalyzer().Flags.Func("exclude", "regexp path (Unix style) to exclude (can be used multiple times)", walkNonEmpty(func(s string) {
		configOverride.Exclude = append(configOverride.Exclude, s)
	}))
	composition.Analyzer.GetAnalyzer().Flags.Func("enable", "comma-separated rule names to enable", walkNonEmptyCSV(func(s string) {
		configOverride.Enable = append(configOverride.Enable, s)
	}))
	composition.Analyzer.GetAnalyzer().Flags.Func("disable", "comma-separated rule names to disable", walkNonEmptyCSV(func(s string) {
		configOverride.Disable = append(configOverride.Disable, s)
	}))

	composition.Analyzer.GetAnalyzer().Flags.BoolFunc("V", "print version and exit", func(s string) error {
		fmt.Println(version.Current)
		os.Exit(0)
		return nil
	})

	singlechecker.Main(composition.Analyzer.GetAnalyzer())
}
