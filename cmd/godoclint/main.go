package main

import (
	"fmt"
	"os"

	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/godoc-lint/godoc-lint/pkg/analysis"
	"github.com/godoc-lint/godoc-lint/pkg/check"
	"github.com/godoc-lint/godoc-lint/pkg/config"
	"github.com/godoc-lint/godoc-lint/pkg/inspect"
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

	analyzer.GetAnalyzer().Flags.BoolFunc("V", "print version and exit", func(s string) error {
		fmt.Println(version.Current)
		os.Exit(0)
		return nil
	})
	singlechecker.Main(analyzer.GetAnalyzer())
}
