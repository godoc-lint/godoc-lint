package package_doc

import (
	"strings"

	"github.com/godoc-lint/godoc-lint/pkg/model"
	"github.com/godoc-lint/godoc-lint/pkg/util"
)

// PackageDocRule is the corresponding rule name.
const PackageDocRule = "package-doc"

var ruleSet = model.RuleSet{}.Add(PackageDocRule)

// PackageDocChecker checks package godocs.
type PackageDocChecker struct{}

// NewPackageDocChecker returns a new instance of the corresponding checker.
func NewPackageDocChecker() *PackageDocChecker {
	return &PackageDocChecker{}
}

// GetCoveredRules implements the corresponding interface method.
func (r *PackageDocChecker) GetCoveredRules() model.RuleSet {
	return ruleSet
}

// Apply implements the corresponding interface method.
func (r *PackageDocChecker) Apply(actx *model.AnalysisContext) error {
	checkPackageDocRule(actx)
	return nil
}

func checkPackageDocRule(actx *model.AnalysisContext) {
	if !actx.Config.IsAnyRuleApplicable(model.RuleSet{}.Add(PackageDocRule)) {
		return
	}

	startWith := strings.TrimSpace(actx.Config.GetRuleOptions().PackageDocStartWith)

	for _, f := range actx.Pass.Files {
		if !util.IsFileApplicable(actx, f) {
			continue
		}

		ir := actx.InspectorResult.Files[f]
		if ir.DisabledRules.All || ir.DisabledRules.Rules.IsSupersetOf(ruleSet) {
			continue
		}

		if ir.PackageDoc == nil || ir.PackageDoc.Text == "" {
			continue
		}

		if ir.PackageDoc.DisabledRules.All || ir.PackageDoc.DisabledRules.Rules.Has(PackageDocRule) {
			continue
		}

		expectedPrefix := f.Name.Name
		if startWith != "" {
			expectedPrefix = startWith + " " + f.Name.Name
		}

		if !strings.HasPrefix(ir.PackageDoc.Text, expectedPrefix) {
			actx.Pass.Reportf(ir.PackageDoc.CG.Pos(), "package godoc should start with %q", expectedPrefix)
		}
	}
}
