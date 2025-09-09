package start_with_name

import (
	"fmt"
	"go/ast"
	"regexp"
	"strings"

	"github.com/godoc-lint/godoc-lint/pkg/check/shared"
	"github.com/godoc-lint/godoc-lint/pkg/model"
	"github.com/godoc-lint/godoc-lint/pkg/util"
)

// StartWithNameRule is the corresponding rule name.
const StartWithNameRule = "start-with-name"

var ruleSet = model.RuleSet{}.Add(StartWithNameRule)

// StartWithNameChecker checks starter of godocs.
type StartWithNameChecker struct{}

// NewStartWithNameChecker returns a new instance of the corresponding checker.
func NewStartWithNameChecker() *StartWithNameChecker {
	return &StartWithNameChecker{}
}

// GetCoveredRules implements the corresponding interface method.
func (r *StartWithNameChecker) GetCoveredRules() model.RuleSet {
	return ruleSet
}

// Apply implements the corresponding interface method.
func (r *StartWithNameChecker) Apply(actx *model.AnalysisContext) error {
	if !actx.Config.IsAnyRuleApplicable(model.RuleSet{}.Add(StartWithNameRule)) {
		return nil
	}

	includeTests := actx.Config.GetRuleOptions().StartWithNameIncludeTests
	includePrivate := actx.Config.GetRuleOptions().StartWithNameIncludeUnexported
	startPattern := actx.Config.GetRuleOptions().StartWithNamePattern
	_, matcher, err := getStartMatcher(startPattern)
	if err != nil {
		return err
	}

	for _, ir := range util.AnalysisApplicableFiles(actx, includeTests, model.RuleSet{}.Add(StartWithNameRule)) {
		for _, decl := range ir.SymbolDecl {
			isExported := ast.IsExported(decl.Name)
			if !isExported && !includePrivate {
				continue
			}

			if decl.Kind == model.SymbolDeclKindBad {
				continue
			}

			if decl.Doc == nil || decl.Doc.Text == "" {
				continue
			}

			if decl.Doc.DisabledRules.All || decl.Doc.DisabledRules.Rules.Has(StartWithNameRule) {
				continue
			}

			if decl.MultiNameDecl {
				continue
			}

			if shared.HasDeprecatedParagraph(decl.Doc.Parsed.Content) {
				// If there's a paragraph starting with "Deprecated:", we skip the
				// entire godoc. The reason is a deprecated symbol will not appear
				// when docs are rendered.
				//
				// Another reason is that we cannot just skip those paragraphs and
				// look for the symbol in the remaining text. To do that, we need
				// to iterate over all comment.Block nodes, and check if a block
				// is a paragraph AND starts with the deprecation marker. This is
				// simple, but the challenge appears when we get to the first block
				// that does not have the marker and we want to check if it starts
				// with the symbol name. We'd expect that to be a paragraph, but
				// that is not always the case. For example, take this decl:
				//
				//    // Deprecated: blah blah
				//    //
				//    // Foo is integer
				//    //
				//    // Deprecation: blah blah
				//    type Foo int
				//
				// The first block is a paragraph which we can easily skip due to
				// the "Deprecated:" marker. However, the second block is actually
				// parsed as a heading. One can verify this by copy/pasting it in
				// a Go file and check the gopls hover.
				//
				// There might be a workaround for this, but this also means the
				// godoc parser behaves in unexpected ways, and until we don't
				// really know the extent of its quirks, it's safer to just skip
				// further checks on such godocs.
				continue
			}

			symbolNameFromDoc := matcher(decl.Doc.Text)
			if symbolNameFromDoc == decl.Name {
				continue
			}

			actx.Pass.ReportRangef(&decl.Doc.CG, "godoc should start with symbol name (%q)", decl.Name)
		}
	}
	return nil
}

const symbolNameSubmatch = "symbol_name"

var symbolNameSubmatchPattern = fmt.Sprintf(`(?P<%s>.+?)\b`, symbolNameSubmatch)

func getStartMatcher(startPattern string) (string, func(string) string, error) {
	var replaced string
	if strings.Contains(startPattern, "%") {
		replaced = strings.ReplaceAll(startPattern, "%", symbolNameSubmatchPattern)
	} else {
		if startPattern == "" || strings.HasSuffix(startPattern, " ") {
			replaced = startPattern + symbolNameSubmatchPattern
		} else {
			replaced = startPattern + " " + symbolNameSubmatchPattern
		}
	}
	if !strings.HasPrefix(replaced, "^") {
		replaced = "^" + replaced
	}

	re, err := regexp.Compile(replaced)
	if err != nil {
		return "", nil, fmt.Errorf("invalid start pattern: %w", err)
	}

	ix := re.SubexpIndex(symbolNameSubmatch)
	if ix == -1 {
		return "", nil, fmt.Errorf("cannot find named group %q in pattern: %q", symbolNameSubmatch, re.String())
	}
	return replaced, func(s string) string {
		match := re.FindStringSubmatch(s)
		if len(match) == 0 {
			return ""
		}
		return match[ix]
	}, nil
}
