package analysis

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"golang.org/x/tools/go/analysis"

	"github.com/godoc-lint/godoc-lint/pkg/model"
	"github.com/godoc-lint/godoc-lint/pkg/version"
)

const (
	metaName = "godoclint"
	metaDoc  = "Checks Golang's documentation practice (godoc)."
	metaURL  = "https://github.com/godoc-lint/godoc-lint"
)

// Analyzer implements the godoc-lint analyzer.
type Analyzer struct {
	cb        model.ConfigBuilder
	inspector model.Inspector
	reg       model.Registry

	analyzer *analysis.Analyzer
}

// NewAnalyzer returns a new instance of the corresponding analyzer.
func NewAnalyzer(cb model.ConfigBuilder, reg model.Registry, inspector model.Inspector) *Analyzer {
	result := &Analyzer{
		cb:        cb,
		reg:       reg,
		inspector: inspector,
		analyzer: &analysis.Analyzer{
			Name:     metaName,
			Doc:      metaDoc,
			URL:      metaURL,
			Requires: []*analysis.Analyzer{inspector.GetAnalyzer()},
		},
	}

	result.analyzer.Flags.BoolFunc("V", "print version and exit", func(s string) error {
		fmt.Printf("%s %s\n", result.analyzer.Name, version.Version)
		os.Exit(0)
		return nil
	})

	configOverride := model.NewConfigOverride()
	cb.SetOverride(configOverride)

	result.analyzer.Flags.Func("config", "path to config file", func(s string) error {
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

	result.analyzer.Flags.Func("include", "regexp path to include (can be used multiple times)", walkNonEmpty(func(s string) {
		configOverride.Include = append(configOverride.Include, s)
	}))
	result.analyzer.Flags.Func("exclude", "regexp path to exclude (can be used multiple times)", walkNonEmpty(func(s string) {
		configOverride.Exclude = append(configOverride.Exclude, s)
	}))
	result.analyzer.Flags.Func("enable", "comma-separated rule names to enable", walkNonEmptyCSV(func(s string) {
		configOverride.Enable = append(configOverride.Enable, s)
	}))
	result.analyzer.Flags.Func("disable", "comma-separated rule names to disable", walkNonEmptyCSV(func(s string) {
		configOverride.Disable = append(configOverride.Disable, s)
	}))

	result.analyzer.Run = result.run
	return result
}

// GetAnalyzer returns the underlying analyzer.
func (a *Analyzer) GetAnalyzer() *analysis.Analyzer {
	return a.analyzer
}

func (a *Analyzer) run(pass *analysis.Pass) (any, error) {
	actx := &model.AnalysisContext{
		Config:          a.cb.MustGetConfig(),
		InspectorResult: pass.ResultOf[a.inspector.GetAnalyzer()].(*model.InspectorResult),
		Pass:            pass,
	}

	for _, rule := range a.reg.Rules() {
		ruleName := rule.GetName()
		if !actx.Config.IsRuleApplicable(ruleName) {
			continue
		}

		if err := rule.Apply(actx); err != nil {
			return nil, fmt.Errorf("error applying rule (%s): %w", ruleName, err)
		}
	}
	return nil, nil
}
