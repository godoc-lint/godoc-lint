package require_doc

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"

	"github.com/godoc-lint/godoc-lint/pkg/model"
	"github.com/godoc-lint/godoc-lint/pkg/util"
)

// RequireDocRule is the corresponding rule name.
const RequireDocRule = "require-doc"

var ruleSet = model.RuleSet{}.Add(RequireDocRule)

// RequireDocChecker checks required godocs.
type RequireDocChecker struct{}

// NewRequireDocChecker returns a new instance of the corresponding checker.
func NewRequireDocChecker() *RequireDocChecker {
	return &RequireDocChecker{}
}

// GetCoveredRules implements the corresponding interface method.
func (r *RequireDocChecker) GetCoveredRules() model.RuleSet {
	return ruleSet
}

// Apply implements the corresponding interface method.
func (r *RequireDocChecker) Apply(actx *model.AnalysisContext) error {
	includeTests := actx.Config.GetRuleOptions().RequireDocIncludeTests
	requirePublic := !actx.Config.GetRuleOptions().RequireDocIgnoreExported
	requirePrivate := !actx.Config.GetRuleOptions().RequireDocIgnoreUnexported

	if !requirePublic && !requirePrivate {
		return nil
	}

	for _, f := range actx.Pass.Files {
		if !util.IsFileApplicable(actx, f) {
			continue
		}

		ft := util.GetPassFileToken(f, actx.Pass)
		if ft == nil {
			continue
		}

		if !includeTests && strings.HasSuffix(ft.Name(), "_test.go") {
			continue
		}

		ir := actx.InspectorResult.Files[f]
		if ir == nil || ir.DisabledRules.All || ir.DisabledRules.Rules.Has(RequireDocRule) {
			continue
		}

		for _, decl := range ir.SymbolDecl {
			isExported := ast.IsExported(decl.Name)
			if isExported && !requirePublic || !isExported && !requirePrivate {
				continue
			}

			if decl.Doc != nil && (decl.Doc.DisabledRules.All || decl.Doc.DisabledRules.Rules.Has(RequireDocRule)) {
				continue
			}

			if decl.Kind == model.SymbolDeclKindBad {
				continue
			}

			if decl.Kind == model.SymbolDeclKindFunc {
				if decl.Doc == nil || decl.Doc.Text == "" {
					reportRange(actx.Pass, decl.Ident)
				}
				continue
			}

			// Now we only have const/var/type declarations.

			if decl.Doc != nil && decl.Doc.Text != "" {
				// cases:
				//
				//   // godoc
				//   const foo = 0
				//
				//   // godoc
				//   const foo, bar = 0, 0
				//
				//   const (
				//       // godoc
				//       foo = 0
				//   )
				//
				//   const (
				//       // godoc
				//       foo, bar = 0, 0
				//   )
				//
				//   // godoc
				//   type foo int
				//
				//   type (
				//       // godoc
				//       foo int
				//   )
				continue
			}

			if decl.TrailingDoc != nil && decl.TrailingDoc.Text != "" {
				// cases:
				//
				//   const foo = 0 // godoc
				//
				//   const foo, bar = 0, 0 // godoc
				//
				//   const (
				//       foo = 0 // godoc
				//   )
				//
				//   const (
				//       foo, bar = 0, 0 // godoc
				//   )
				//
				//   type foo int // godoc
				//
				//   type (
				//       foo int // godoc
				//   )
				continue
			}

			if decl.ParentDoc != nil && decl.ParentDoc.Text != "" {
				// cases:
				//
				//   // godoc
				//   const (
				//       foo = 0
				//   )
				//
				//   // godoc
				//   const (
				//       foo, bar = 0, 0
				//   )
				//
				//   // godoc
				//   type (
				//       foo int
				//   )
				continue
			}

			// At this point there is no godoc for the symbol.
			//
			// cases:
			//
			//   const foo = 0
			//
			//   const foo, bar = 0, 0
			//
			//   const (
			//       foo = 0
			//   )
			//
			//   const (
			//       foo, bar = 0, 0
			//   )
			//
			//   type foo int
			//
			//   type (
			//       foo int
			//   )

			reportRange(actx.Pass, decl.Ident)
		}
	}
	return nil
}

func reportRange(pass *analysis.Pass, ident *ast.Ident) {
	pass.ReportRangef(ident, "symbol should have a godoc (%v)", ident.Name)
}
