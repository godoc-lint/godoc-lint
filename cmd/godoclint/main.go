package main

import (
	"fmt"
	"os"

	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/godoc-lint/godoc-lint/pkg/analysis"
	"github.com/godoc-lint/godoc-lint/pkg/config"
	"github.com/godoc-lint/godoc-lint/pkg/inspect"
	"github.com/godoc-lint/godoc-lint/pkg/rule"
)

func main() {
	exitFunc := func(code int, err error) {
		fmt.Fprintf(os.Stderr, "%v\n", err.Error())
		os.Exit(code)
	}
	reg := rule.NewPopulatedRegistry()
	cb := config.NewConfigBuilder("", reg.GetCoveredRules())
	ocb := config.NewOnceConfigBuilder(cb)
	inspector := inspect.NewInspector(ocb, exitFunc)
	analyzer := analysis.NewAnalyzer(ocb, reg, inspector, exitFunc)
	singlechecker.Main(analyzer.GetAnalyzer())
}
