package rule

import (
	"go/token"

	"golang.org/x/tools/go/analysis"

	"github.com/godoc-lint/godoc-lint/pkg/model"
)

// MaxLengthRuleName is the corresponding rule name/identifier.
const MaxLengthRuleName = "max-length"

// MaxLengthRule is a rule that checks maximum line length of godocs.
type MaxLengthRule struct{}

// NewMaxLengthRule returns a new instance of the corresponding rule.
func NewMaxLengthRule() *MaxLengthRule {
	return &MaxLengthRule{}
}

// GetName returns the name of the rule.
func (r *MaxLengthRule) GetName() string {
	return MaxLengthRuleName
}

// Apply checks for the rule.
func (r *MaxLengthRule) Apply(actx *model.AnalysisContext, pass *analysis.Pass) error {
	for _, f := range pass.Files {
		if f.Pos() == token.NoPos {
			continue
		}
		ft := pass.Fset.File(f.Pos())
		if ft == nil {
			continue
		}
		if !actx.Config.IsPathApplicable(ft.Name()) {
			continue
		}
	}
	return nil
}
