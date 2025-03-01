package package_doc

import (
	"go/ast"
	"strings"

	"github.com/godoc-lint/godoc-lint/pkg/model"
	"github.com/godoc-lint/godoc-lint/pkg/util"
)

const (
	PkgDocRule       = "pkg-doc"
	SinglePkgDocRule = "single-pkg-doc"
)

var ruleSet = model.RuleSet{}.Add(PkgDocRule, SinglePkgDocRule)

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
	checkSinglePackageDocRule(actx)
	return nil
}

func checkPackageDocRule(actx *model.AnalysisContext) {
	if !actx.Config.IsAnyRuleApplicable(model.RuleSet{}.Add(PkgDocRule)) {
		return
	}

	startWith := strings.TrimSpace(actx.Config.GetRuleOptions().PkgDocStartWith)

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

		if ir.PackageDoc.DisabledRules.All || ir.PackageDoc.DisabledRules.Rules.Has(PkgDocRule) {
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

func checkSinglePackageDocRule(actx *model.AnalysisContext) {
	if !actx.Config.IsAnyRuleApplicable(model.RuleSet{}.Add(SinglePkgDocRule)) {
		return
	}

	documentedPkgs := make(map[string][]*ast.File, 2)

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

		if ir.PackageDoc.DisabledRules.All || ir.PackageDoc.DisabledRules.Rules.Has(SinglePkgDocRule) {
			continue
		}

		pkg := f.Name.Name
		if _, ok := documentedPkgs[pkg]; !ok {
			documentedPkgs[pkg] = make([]*ast.File, 0, 2)
		}
		documentedPkgs[pkg] = append(documentedPkgs[pkg], f)
	}

	for _, fs := range documentedPkgs {
		if len(fs) < 2 {
			continue
		}
		for _, f := range fs {
			ir := actx.InspectorResult.Files[f]
			actx.Pass.Reportf(ir.PackageDoc.CG.Pos(), "package should have a single godoc (%d found)", len(fs))
		}
	}
}
