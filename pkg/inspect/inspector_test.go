package inspect_test

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/tools/go/analysis/analysistest"
	"gopkg.in/yaml.v3"

	"github.com/godoc-lint/godoc-lint/pkg/config"
	"github.com/godoc-lint/godoc-lint/pkg/inspect"
	"github.com/godoc-lint/godoc-lint/pkg/model"
	"github.com/godoc-lint/godoc-lint/pkg/rule"
	"github.com/godoc-lint/godoc-lint/pkg/util"
)

var _ model.Inspector = &inspect.Inspector{}

func TestInspector(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	wd, err := os.Getwd()
	require.Nil(err, "failed to get wd")

	exitFunc := func(code int) {}

	testdir := filepath.Join(wd, "../../testdata/inspector")

	reg := rule.NewPopulatedRegistry()
	cb := config.NewConfigBuilder(testdir, reg.GetCoveredRules(), exitFunc)
	inspector := inspect.NewInspector(cb)

	ars := analysistest.Run(t, testdir, inspector.GetAnalyzer(), "./...")

	for _, ar := range ars {
		result, ok := ar.Result.(*model.InspectorResult)
		require.True(ok, "unknown result type")

		for _, f := range ar.Pass.Files {
			ft := util.GetPassFileToken(f, ar.Pass)
			refFile := strings.Replace(ft.Name(), filepath.Ext(ft.Name()), ".yaml", 1)
			_, statErr := os.Stat(refFile)

			resultEntry, ok := result.Files[f]
			require.False(!ok && statErr == nil, "unexpected test-ref file for %q", ft.Name())
			require.False(ok && statErr != nil, "test-ref file not found for %q", ft.Name())

			if !ok {
				continue
			}

			buf := bytes.NewBuffer(nil)
			enc := yaml.NewEncoder(buf)
			enc.SetIndent(2)

			got := simplifyResultEntry(resultEntry)
			err := enc.Encode(got)
			require.Nil(err, "cannot marshal got value: %v", err)
			rawGot := buf.Bytes()

			rawRef, err := os.ReadFile(refFile)
			require.Nil(err, "cannot read ref file %q: %v", refFile, err)

			match := assert.YAMLEqf(string(rawRef), string(rawGot), "got and ref do not match for %q", refFile)
			gotFile := strings.Replace(refFile, ".yaml", ".got.yaml", 1)
			if !match {
				t.Logf("writing got value to %q", gotFile)
				err := os.WriteFile(gotFile, rawGot, os.ModePerm)
				assert.Nil(err, "cannot write got value to file %q", gotFile)
			} else {
				_ = os.Remove(gotFile)
			}
		}
	}
}

func simplifyResultEntry(entry *model.FileInspection) any {
	doc := func(doc *model.CommentGroup) *string {
		if doc == nil {
			return nil
		}
		text := doc.CG.Text()
		return &text
	}

	m := map[string]any{
		"disabled-rules": map[string]any{
			"all":   entry.DisabledRules.All,
			"rules": entry.DisabledRules.Rules.List(),
		},
		"package-doc": doc(entry.PackageDoc),
	}
	if entry.SymbolDecl != nil {
		sds := make([]any, 0, len(entry.SymbolDecl))
		for _, sd := range entry.SymbolDecl {
			sds = append(sds, map[string]any{
				"kind":            sd.Kind,
				"name":            sd.Name,
				"is-type-alias":   sd.IsTypeAlias,
				"multi-name-decl": sd.MultiNameDecl,
				"parent-doc":      doc(sd.ParentDoc),
				"doc":             doc(sd.Doc),
				"disabled-rules": map[string]any{
					"all":   sd.DisabledRules.All,
					"rules": sd.DisabledRules.Rules.List(),
				},
			})
		}
		m["symbol-decl"] = sds
	}
	return m
}
