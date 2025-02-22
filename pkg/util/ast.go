package util

import (
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"

	"github.com/godoc-lint/godoc-lint/pkg/model"
)

// GetPassFileToken is a helper function to return the file token associated
// with the given AST file.
func GetPassFileToken(f *ast.File, pass *analysis.Pass) *token.File {
	if f.Pos() == token.NoPos {
		return nil
	}
	ft := pass.Fset.File(f.Pos())
	if ft == nil {
		return nil
	}
	return ft
}

// IsFileApplicable is a helper function to determine if the given AST file
// should be included in analysis.
func IsFileApplicable(actx *model.AnalysisContext, f *ast.File) bool {
	if actx.InspectorResult == nil || actx.InspectorResult.Files[f] == nil {
		return false
	}
	ft := GetPassFileToken(f, actx.Pass)
	if ft == nil {
		return false
	}
	return actx.Config.IsPathApplicable(ft.Name())
}
