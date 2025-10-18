package util

import (
	"go/ast"
	"go/token"
	"iter"
	"strings"

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

// AnalysisApplicableFiles returns an iterator looping over files that are ready
// to be analyzed.
//
// The yield-ed arguments are never nil.
func AnalysisApplicableFiles(actx *model.AnalysisContext, includeTests bool, ruleSet model.RuleSet) iter.Seq2[*ast.File, *model.FileInspection] {
	return func(yield func(*ast.File, *model.FileInspection) bool) {
		if actx.InspectorResult == nil {
			return
		}

		for _, f := range actx.Pass.Files {
			ir := actx.InspectorResult.Files[f]

			if ir == nil {
				continue
			}

			ft := GetPassFileToken(f, actx.Pass)
			if ft == nil {
				continue
			}

			if !actx.Config.IsPathApplicable(ft.Name()) {
				continue
			}

			if !includeTests && strings.HasSuffix(ft.Name(), "_test.go") {
				continue
			}

			if ir.DisabledRules.All || ir.DisabledRules.Rules.IsSupersetOf(ruleSet) {
				continue
			}

			if !yield(f, ir) {
				return
			}
		}
	}
}

// IsMethodOnUnexportedReceiver checks if a declaration is a method on an unexported type.
// It returns true if the declaration is a method AND its receiver type is unexported.
func IsMethodOnUnexportedReceiver(decl ast.Decl) bool {
	// Only function declarations can be methods
	funcDecl, ok := decl.(*ast.FuncDecl)
	if !ok {
		return false
	}

	// If there's no receiver, it's a function, not a method
	if funcDecl.Recv == nil || len(funcDecl.Recv.List) == 0 {
		return false
	}

	// Extract the receiver type name
	receiverType := funcDecl.Recv.List[0].Type

	// Get the type identifier, handling both direct types and pointer types
	var typeIdent *ast.Ident
	switch t := receiverType.(type) {
	case *ast.Ident:
		// Direct receiver: func (r MyType) Method()
		typeIdent = t
	case *ast.StarExpr:
		// Pointer receiver: func (r *MyType) Method()
		// Also handles pointer to generic types: func (r *MyType[T]) Method() or func (r *MyType[T, U]) Method()
		switch x := t.X.(type) {
		case *ast.Ident:
			typeIdent = x
		case *ast.IndexExpr:
			// Pointer to generic type with one parameter: func (r *MyType[T]) Method()
			if ident, ok := x.X.(*ast.Ident); ok {
				typeIdent = ident
			}
		case *ast.IndexListExpr:
			// Pointer to generic type with multiple parameters: func (r *MyType[T, U]) Method()
			if ident, ok := x.X.(*ast.Ident); ok {
				typeIdent = ident
			}
		}
	case *ast.IndexExpr:
		// Generic type with one parameter: func (r MyType[T]) Method()
		if ident, ok := t.X.(*ast.Ident); ok {
			typeIdent = ident
		}
	case *ast.IndexListExpr:
		// Generic type with multiple parameters: func (r MyType[T, U]) Method()
		if ident, ok := t.X.(*ast.Ident); ok {
			typeIdent = ident
		}
	}

	// If we found a type identifier, check if it's unexported
	if typeIdent != nil {
		return !ast.IsExported(typeIdent.Name)
	}

	return false
}
