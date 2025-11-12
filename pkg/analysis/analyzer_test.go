package analysis_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/godoc-lint/godoc-lint/pkg/analysis"
	"github.com/godoc-lint/godoc-lint/pkg/check"
	"github.com/godoc-lint/godoc-lint/pkg/config"
	"github.com/godoc-lint/godoc-lint/pkg/inspect"
)

func TestRules(t *testing.T) {
	require := require.New(t)

	wd, err := os.Getwd()
	require.NoError(err, "failed to get wd")

	exitFunc := func(code int, err error) {
		panic(fmt.Sprintf("exit code %d: %v", code, err))
	}

	testdir := filepath.Join(wd, "../../testdata/rule")

	reg := check.NewPopulatedRegistry()
	cb := config.NewConfigBuilder(testdir)
	ocb := config.NewOnceConfigBuilder(cb)
	inspector := inspect.NewInspector(ocb, exitFunc)
	analyzer := analysis.NewAnalyzer(testdir, ocb, reg, inspector, exitFunc)

	_ = analysistest.Run(t, testdir, analyzer.GetAnalyzer(), "./...")
}
