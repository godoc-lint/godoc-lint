package pkg_doc

import (
	"go/ast"
	"strings"

	"github.com/godoc-lint/godoc-lint/pkg/model"
	"github.com/godoc-lint/godoc-lint/pkg/util"
)

const (
	PkgDocRule        = "pkg-doc"
	SinglePkgDocRule  = "single-pkg-doc"
	RequirePkgDocRule = "require-pkg-doc"
)

var ruleSet = model.RuleSet{}.Add(
	PkgDocRule,
	SinglePkgDocRule,
	RequirePkgDocRule,
)

// PkgDocChecker checks package godocs.
type PkgDocChecker struct{}

// NewPkgDocChecker returns a new instance of the corresponding checker.
func NewPkgDocChecker() *PkgDocChecker {
	return &PkgDocChecker{}
}

// GetCoveredRules implements the corresponding interface method.
func (r *PkgDocChecker) GetCoveredRules() model.RuleSet {
	return ruleSet
}

// Apply implements the corresponding interface method.
func (r *PkgDocChecker) Apply(actx *model.AnalysisContext) error {
	checkPkgDocRule(actx)
	checkSinglePkgDocRule(actx)
	checkRequirePkgDocRule(actx)
	return nil
}

const commandPkgName = "main"
const commandTestPkgName = "main_test"

func checkPkgDocRule(actx *model.AnalysisContext) {
	if !actx.Config.IsAnyRuleApplicable(model.RuleSet{}.Add(PkgDocRule)) {
		return
	}

	includeTests := actx.Config.GetRuleOptions().PkgDocIncludeTests
	startWith := strings.TrimSpace(actx.Config.GetRuleOptions().PkgDocStartWith)

	for f, ir := range util.AnalysisApplicableFiles(actx, includeTests, model.RuleSet{}.Add(PkgDocRule)) {
		if ir.PackageDoc == nil {
			continue
		}

		if ir.PackageDoc.DisabledRules.All || ir.PackageDoc.DisabledRules.Rules.Has(PkgDocRule) {
			continue
		}

		if f.Name.Name == commandPkgName || f.Name.Name == commandTestPkgName {
			// Skip command packages, as they are not required to start with
			// "Package main" or "Package main_test".
			//
			// See for more details:
			//   - https://github.com/godoc-lint/godoc-lint/issues/10
			//   - https://go.dev/doc/comment#cmd
			continue
		}

		if expectedPrefix, ok := checkPkgDocPrefix(ir.PackageDoc.Text, startWith, f.Name.Name); !ok {
			actx.Pass.Reportf(ir.PackageDoc.CG.Pos(), "package godoc should start with %q", expectedPrefix+" ")
		}
	}
}

func checkPkgDocPrefix(text string, startWith string, packageName string) (string, bool) {
	if text == "" {
		return "", true
	}
	expectedPrefix := packageName
	if startWith != "" {
		expectedPrefix = startWith + " " + packageName
	}
	if !strings.HasPrefix(text, expectedPrefix) {
		return expectedPrefix, false
	}
	rest := text[len(expectedPrefix):]
	return expectedPrefix, rest == "" || rest[0] == ' ' || rest[0] == '\t' || rest[0] == '\r' || rest[0] == '\n'
}

func checkSinglePkgDocRule(actx *model.AnalysisContext) {
	if !actx.Config.IsAnyRuleApplicable(model.RuleSet{}.Add(SinglePkgDocRule)) {
		return
	}

	includeTests := actx.Config.GetRuleOptions().SinglePkgDocIncludeTests

	documentedPkgs := make(map[string][]*ast.File, 2)

	for f, ir := range util.AnalysisApplicableFiles(actx, includeTests, model.RuleSet{}.Add(SinglePkgDocRule)) {
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

func checkRequirePkgDocRule(actx *model.AnalysisContext) {
	if !actx.Config.IsAnyRuleApplicable(model.RuleSet{}.Add(RequirePkgDocRule)) {
		return
	}

	includeTests := actx.Config.GetRuleOptions().RequirePkgDocIncludeTests

	pkgFiles := make(map[string][]*ast.File, 2)

	for f := range util.AnalysisApplicableFiles(actx, includeTests, model.RuleSet{}.Add(RequirePkgDocRule)) {
		pkg := f.Name.Name
		if _, ok := pkgFiles[pkg]; !ok {
			pkgFiles[pkg] = make([]*ast.File, 0, len(actx.Pass.Files))
		}
		pkgFiles[pkg] = append(pkgFiles[pkg], f)
	}

	for _, fs := range pkgFiles {
		pkgHasGodoc := false
		for _, f := range fs {
			ir := actx.InspectorResult.Files[f]

			if ir.PackageDoc == nil || ir.PackageDoc.Text == "" {
				continue
			}

			if ir.PackageDoc.DisabledRules.All || ir.PackageDoc.DisabledRules.Rules.Has(RequirePkgDocRule) {
				continue
			}

			pkgHasGodoc = true
			break
		}

		if pkgHasGodoc {
			continue
		}

		// Add a diagnostic message to the first file of the package.
		actx.Pass.Reportf(fs[0].Name.Pos(), "package should have a godoc")
	}
}
