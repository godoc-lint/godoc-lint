package max_length

import (
	"fmt"
	gdc "go/doc/comment"
	"strings"

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
	ro := actx.Config.GetRuleOptions()
	if ro == nil || ro.MaxLength == nil {
		panic("missing rule options")
	}
	maxLength := int(ro.MaxLength.Length)

	for _, f := range actx.Pass.Files {
		if !util.IsFileApplicable(actx, f) {
			continue
		}

		ir := actx.InspectorResult.Files[f]
		if ir.DisabledRules.All {
			continue
		}
		if _, ok := ir.DisabledRules.Rules[MaxLengthRuleName]; ok {
			continue
		}

		if ir.PackageDoc != nil {
			checkMaxLength(actx, ir.PackageDoc, maxLength)
		}

		processedParents := make(map[*model.CommentGroup]struct{}, len(ir.SymbolDecl))
		for _, sd := range ir.SymbolDecl {
			if sd.DisabledRules.All {
				continue
			}
			if _, ok := sd.DisabledRules.Rules[MaxLengthRuleName]; ok {
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
	docWithoutCodeBlocks := &gdc.Doc{
		Content: nonCodeBlocks,
		Links:   doc.Parsed.Links,
	}
	text := string((&gdc.Printer{}).Comment(docWithoutCodeBlocks))
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
