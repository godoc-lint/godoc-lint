package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/godoc-lint/godoc-lint/pkg/analysis"
	"github.com/godoc-lint/godoc-lint/pkg/config"
	"github.com/godoc-lint/godoc-lint/pkg/inspect"
	"github.com/godoc-lint/godoc-lint/pkg/rule"
)

func main() {
	reg := rule.NewPopulatedRegistry()
	cb := config.NewConfigBuilder("", reg.GetCoveredRules(), nil)
	inspector := inspect.NewInspector(cb)
	analyzer := analysis.NewAnalyzer(cb, reg, inspector)
	singlechecker.Main(analyzer.GetAnalyzer())
}
