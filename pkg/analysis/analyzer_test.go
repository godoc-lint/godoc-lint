package analysis_test

import (
	"os"
	"path/filepath"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/godoc-lint/godoc-lint/pkg/analysis"
	"github.com/godoc-lint/godoc-lint/pkg/config"
	"github.com/godoc-lint/godoc-lint/pkg/inspect"
	"github.com/godoc-lint/godoc-lint/pkg/rule"
)

func TestTypeGodoc(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get wd: %s", err)
	}

	exitFunc := func(code int) {}

	reg := rule.NewPopulatedRegistry()
	cb := config.NewConfigBuilder("", reg.Names(), exitFunc)
	inspector := inspect.NewInspector(cb)
	analyzer := analysis.NewAnalyzer(cb, reg, inspector)

	testdata := filepath.Join(wd, "../../testdata/type-godoc")
	analysistest.Run(
		&testing.T{},
		testdata,
		analyzer.GetAnalyzer(),
		"",
	)
}
