package max_length

import (
	"fmt"
	gdc "go/doc/comment"
	"strings"

	"github.com/godoc-lint/godoc-lint/pkg/model"
	"github.com/godoc-lint/godoc-lint/pkg/util"
)

// MaxLengthRule is the corresponding rule name.
const MaxLengthRule = "max-length"

var ruleSet = model.RuleSet{}.Add(MaxLengthRule)

// MaxLengthChecker checks maximum line length of godocs.
type MaxLengthChecker struct{}

// NewMaxLengthChecker returns a new instance of the corresponding checker.
func NewMaxLengthChecker() *MaxLengthChecker {
	return &MaxLengthChecker{}
}

// GetCoveredRules implements the corresponding interface method.
func (r *MaxLengthChecker) GetCoveredRules() model.RuleSet {
	return ruleSet
}

// Apply implements the corresponding interface method.
func (r *MaxLengthChecker) Apply(actx *model.AnalysisContext) error {
	maxLength := int(actx.Config.GetRuleOptions().MaxLength)

	for _, f := range actx.Pass.Files {
		if !util.IsFileApplicable(actx, f) {
			continue
		}

		ir := actx.InspectorResult.Files[f]
		if ir.DisabledRules.All || ir.DisabledRules.Rules.IsSupersetOf(ruleSet) {
			continue
		}

		if ir.PackageDoc != nil {
			checkMaxLength(actx, ir.PackageDoc, maxLength)
		}

		processedParents := make(map[*model.CommentGroup]struct{}, len(ir.SymbolDecl))
		for _, sd := range ir.SymbolDecl {
			if sd.DisabledRules.All || sd.DisabledRules.Rules.Has(MaxLengthRule) {
				continue
			}

			if sd.ParentDoc != nil {
				if _, ok := processedParents[sd.ParentDoc]; !ok {
					processedParents[sd.ParentDoc] = struct{}{}
					checkMaxLength(actx, sd.ParentDoc, maxLength)
				}
			}
			if sd.Doc == nil {
				continue
			}
			checkMaxLength(actx, sd.Doc, maxLength)
		}
	}
	return nil
}

func checkMaxLength(actx *model.AnalysisContext, doc *model.CommentGroup, maxLength int) {
	linkDefsMap := make(map[string]struct{}, len(doc.Parsed.Links))
	for _, linkDef := range doc.Parsed.Links {
		linkDefLine := fmt.Sprintf("[%s]: %s", linkDef.Text, linkDef.URL)
		linkDefsMap[linkDefLine] = struct{}{}
	}

	nonCodeBlocks := make([]gdc.Block, 0, len(doc.Parsed.Content))
	for _, b := range doc.Parsed.Content {
		if _, ok := b.(*gdc.Code); ok {
			continue
		}
		nonCodeBlocks = append(nonCodeBlocks, b)
	}
	strippedCodeAndLinks := &gdc.Doc{
		Content: nonCodeBlocks,
	}
	text := string((&gdc.Printer{}).Comment(strippedCodeAndLinks))
	lines := strings.Split(removeCarriageReturn(text), "\n")

	for _, l := range lines {
		if len(l) <= maxLength {
			continue
		}
		actx.Pass.ReportRangef(&doc.CG, "godoc exceeds max length (%d > %d)", len(l), maxLength)
		break
	}
}

func removeCarriageReturn(s string) string {
	return strings.ReplaceAll(s, "\r", "")
}
