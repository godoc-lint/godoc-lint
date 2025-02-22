package rule

import (
	"github.com/godoc-lint/godoc-lint/pkg/model"
	"github.com/godoc-lint/godoc-lint/pkg/util"
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
func (r *MaxLengthRule) Apply(actx *model.AnalysisContext) error {
	for _, f := range actx.Pass.Files {
		if !util.IsFileApplicable(actx, f) {
			continue
		}
	}
	return nil
}
