package main

import (
	"fmt"
	"os"

	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/godoc-lint/godoc-lint/pkg/analysis"
	"github.com/godoc-lint/godoc-lint/pkg/check"
	"github.com/godoc-lint/godoc-lint/pkg/config"
	"github.com/godoc-lint/godoc-lint/pkg/inspect"
)

func main() {
	exitFunc := func(code int, err error) {
		fmt.Fprintf(os.Stderr, "%v\n", err.Error())
		os.Exit(code)
	}

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to get current working directory: %v\n", err)
		os.Exit(1)
	}

	reg := check.NewPopulatedRegistry()
	cb := config.NewConfigBuilder(cwd, reg.GetCoveredRules())
	ocb := config.NewOnceConfigBuilder(cb)
	inspector := inspect.NewInspector(ocb, exitFunc)
	analyzer := analysis.NewAnalyzer(ocb, reg, inspector, exitFunc)
	singlechecker.Main(analyzer.GetAnalyzer())
}
