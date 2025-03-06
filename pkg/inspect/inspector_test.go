package inspect_test

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/tools/go/analysis/analysistest"
	"gopkg.in/yaml.v3"

	"github.com/godoc-lint/godoc-lint/pkg/check"
	"github.com/godoc-lint/godoc-lint/pkg/config"
	"github.com/godoc-lint/godoc-lint/pkg/inspect"
	"github.com/godoc-lint/godoc-lint/pkg/model"
	"github.com/godoc-lint/godoc-lint/pkg/util"
)

var _ model.Inspector = &inspect.Inspector{}

func TestInspector(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	wd, err := os.Getwd()
	require.Nil(err, "failed to get wd")

	exitFunc := func(code int, err error) {
		panic(fmt.Sprintf("exit code %d: %v", code, err))
	}

	testdir := filepath.Join(wd, "../../testdata/inspector")

	reg := check.NewPopulatedRegistry()
	cb := config.NewConfigBuilder(testdir, reg.GetCoveredRules())
	ocb := config.NewOnceConfigBuilder(cb)
	inspector := inspect.NewInspector(ocb, exitFunc)

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
	disabledRules := func(dr model.InspectorResultDisableRules) map[string]any {
		m := map[string]any{}
		if dr.All {
			m["all"] = true
		}
		if len(dr.Rules.List()) > 0 {
			m["rules"] = dr.Rules.List()
		}
		if len(m) == 0 {
			return nil
		}
		return m
	}

	doc := func(doc *model.CommentGroup) map[string]any {
		if doc == nil {
			return nil
		}
		m := map[string]any{}
		if doc.Text != "" {
			m["text"] = doc.Text
		}
		if subm := disabledRules(doc.DisabledRules); subm != nil {
			m["disabled-rules"] = subm
		}
		if len(m) == 0 {
			return nil
		}
		return m
	}

	m := map[string]any{}
	if subm := disabledRules(entry.DisabledRules); subm != nil {
		m["disabled-rules"] = subm
	}
	if subm := doc(entry.PackageDoc); subm != nil {
		m["package-doc"] = subm
	}
	if entry.SymbolDecl != nil {
		sds := make([]any, 0, len(entry.SymbolDecl))
		for _, sd := range entry.SymbolDecl {
			item := map[string]any{
				"kind": sd.Kind,
				"name": sd.Name,
			}
			if sd.IsTypeAlias {
				item["is-type-alias"] = sd.IsTypeAlias
			}
			if sd.MultiNameDecl {
				item["multi-name-decl"] = sd.MultiNameDecl
			}
			if subm := doc(sd.ParentDoc); subm != nil {
				item["parent-doc"] = subm
			}
			if subm := doc(sd.Doc); subm != nil {
				item["doc"] = subm
			}
			sds = append(sds, item)
		}
		if len(sds) > 0 {
			m["symbol-decl"] = sds
		}
	}
	return m
}
