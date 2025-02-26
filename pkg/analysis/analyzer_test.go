package analysis_test

import (
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/godoc-lint/godoc-lint/pkg/analysis"
	"github.com/godoc-lint/godoc-lint/pkg/config"
	"github.com/godoc-lint/godoc-lint/pkg/inspect"
	"github.com/godoc-lint/godoc-lint/pkg/rule"
)

func TestRules(t *testing.T) {
	require := require.New(t)

	wd, err := os.Getwd()
	require.Nil(err, "failed to get wd")

	basedir := filepath.Join(wd, "../../testdata/rule")

	type testdir struct {
		configDir string
		path      string
	}
	testdirs := []testdir{}

	const configFileName = ".godoc-lint.yaml"

	var walk func(path string, lastConfigDir string)
	walk = func(path string, lastConfigDir string) {
		entries, err := os.ReadDir(path)
		require.Nil(err, "cannot read dir")

		hasConfigFile := slices.ContainsFunc(entries, func(fi os.DirEntry) bool {
			return !fi.IsDir() && fi.Name() == configFileName
		})
		if hasConfigFile {
			lastConfigDir = path
		}

		hasGoFiles := slices.ContainsFunc(entries, func(fi os.DirEntry) bool {
			return !fi.IsDir() && strings.HasSuffix(fi.Name(), ".go")
		})
		if hasGoFiles {
			require.NotEmpty(lastConfigDir, "no config file in (or in parent of) %q", path)
			testdirs = append(testdirs, testdir{
				configDir: lastConfigDir,
				path:      path,
			})
		}

		for _, entry := range entries {
			if !entry.IsDir() {
				continue
			}
			walk(filepath.Join(path, entry.Name()), lastConfigDir)
		}
	}

	walk(basedir, "")

	for _, td := range testdirs {
		relativePath, err := filepath.Rel(basedir, td.path)
		require.Nil(err, "cannot convert to relative path")

		t.Run(relativePath, func(t *testing.T) {
			exitFunc := func(code int) {}

			reg := rule.NewPopulatedRegistry()
			cb := config.NewConfigBuilder(td.configDir, reg.Names(), exitFunc)
			inspector := inspect.NewInspector(cb)
			analyzer := analysis.NewAnalyzer(cb, reg, inspector)

			analysistest.Run(t, td.path, analyzer.GetAnalyzer(), "./")
		})
	}
}
