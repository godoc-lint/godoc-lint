package stdlib_doclink

import (
	"fmt"
	gdc "go/doc/comment"
	"regexp"
	"slices"
	"strings"
	"sync"

	"github.com/godoc-lint/godoc-lint/pkg/check/stdlib_doclink/internal"
	"github.com/godoc-lint/godoc-lint/pkg/model"
	"github.com/godoc-lint/godoc-lint/pkg/util"
)

// StdlibDoclinkRule is the corresponding rule name.
const StdlibDoclinkRule = model.StdlibDoclinkRule

var ruleSet = model.RuleSet{}.Add(StdlibDoclinkRule)

// StdlibDoclinkChecker checks for proper doc links to stdlib symbols.
type StdlibDoclinkChecker struct{}

// NewStdlibDoclinkChecker returns a new instance of the corresponding checker.
func NewStdlibDoclinkChecker() *StdlibDoclinkChecker {
	return &StdlibDoclinkChecker{}
}

// GetCoveredRules implements the corresponding interface method.
func (r *StdlibDoclinkChecker) GetCoveredRules() model.RuleSet {
	return ruleSet
}

// Apply implements the corresponding interface method.
func (r *StdlibDoclinkChecker) Apply(actx *model.AnalysisContext) error {
	includeTests := actx.Config.GetRuleOptions().StdlibDoclinkIncludeTests

	docs := make(map[*model.CommentGroup]struct{}, 10*len(actx.InspectorResult.Files))

	for _, ir := range util.AnalysisApplicableFiles(actx, includeTests, model.RuleSet{}.Add(StdlibDoclinkRule)) {
		if ir.PackageDoc != nil {
			docs[ir.PackageDoc] = struct{}{}
		}

		for _, sd := range ir.SymbolDecl {
			if sd.ParentDoc != nil {
				docs[sd.ParentDoc] = struct{}{}
			}
			if sd.Doc == nil {
				continue
			}
			docs[sd.Doc] = struct{}{}
		}
	}

	for doc := range docs {
		checkStdlibDoclink(actx, doc)
	}
	return nil
}

func checkStdlibDoclink(actx *model.AnalysisContext, doc *model.CommentGroup) {
	if doc.DisabledRules.All || doc.DisabledRules.Rules.Has(StdlibDoclinkRule) {
		return
	}

	applicableBlocks := make([]gdc.Block, 0, len(doc.Parsed.Content))
	for _, b := range doc.Parsed.Content {
		if _, ok := b.(*gdc.Code); ok {
			continue
		}
		// Doc links are not picked up in headings.
		if _, ok := b.(*gdc.Heading); ok {
			continue
		}
		applicableBlocks = append(applicableBlocks, b)
	}
	strippedCodeAndLinks := &gdc.Doc{
		Content: applicableBlocks,
	}
	text := string((&gdc.Printer{}).Comment(strippedCodeAndLinks))

	for _, pd := range findPotentialDoclinks(text) {
		var founds string
		if pd.count > 1 {
			founds = fmt.Sprintf(" (%d instances)", pd.count)
		}
		if pd.kind == internal.SymbolKindNA {
			actx.Pass.ReportRangef(&doc.CG, "text %q should be replaced with %q to link to stdlib package%s", pd.original, pd.doclink, founds)
		} else {
			actx.Pass.ReportRangef(&doc.CG, "text %q should be replaced with %q to link to stdlib %s%s", pd.original, pd.doclink, kindTitle(pd.kind), founds)
		}
	}
}

func kindTitle(kind internal.SymbolKind) string {
	switch kind {
	case internal.SymbolKindType:
		return "type"
	case internal.SymbolKindFunc:
		return "function"
	case internal.SymbolKindVar:
		return "variable"
	case internal.SymbolKindConst:
		return "constant"
	case internal.SymbolKindMethod:
		return "method"
	default:
		return "symbol"
	}
}

type potentialDoclink struct {
	original string
	count    int
	doclink  string
	kind     internal.SymbolKind
}

var potentialDoclinkRE = sync.OnceValue(func() *regexp.Regexp {
	quotedPkgPaths := make([]string, 0, len(stdlib()))
	for _, s := range stdlib() {
		quotedPkgPaths = append(quotedPkgPaths, regexp.QuoteMeta(s.Path))
	}

	// The package-only case is not a match due to potential false positives matching
	// common words like "bytes", or "time".
	stdlibPkgRE, _ := regexp.Compile(fmt.Sprintf(`(?m)(?:^| )(%s)\.([a-zA-Z0-9_]+)(?:\.([a-zA-Z0-9_]+))?\b`, strings.Join(quotedPkgPaths, "|")))
	// Error is never nil due to tests.
	return stdlibPkgRE
})

func findPotentialDoclinks(text string) []potentialDoclink {
	m := make(map[string]*potentialDoclink, 5)

	matches := potentialDoclinkRE().FindAllStringSubmatch(text, -1)
	for _, match := range matches {
		pkg := match[1]
		name1 := match[2]
		name2 := match[3]

		if pkg != "" && name1 != "" && name2 != "" {
			// pkg/name1.name2
			if s, ok := stdlib()[pkg]; ok {
				if kind, ok := s.Symbols[name1+"."+name2]; ok {
					original := fmt.Sprintf("%s.%s.%s", pkg, name1, name2)
					if _, ok := m[original]; !ok {
						doclink := fmt.Sprintf("[%s]", original)
						m[original] = &potentialDoclink{original: original, doclink: doclink, kind: kind}
					}
					m[original].count = m[original].count + 1
				}
			}
		} else if pkg != "" && name1 != "" && name2 == "" {
			// pkg/name1
			if s, ok := stdlib()[pkg]; ok {
				if kind, ok := s.Symbols[name1]; ok {
					original := fmt.Sprintf("%s.%s", pkg, name1)
					if _, ok := m[original]; !ok {
						doclink := fmt.Sprintf("[%s]", original)
						m[original] = &potentialDoclink{original: original, doclink: doclink, kind: kind}
					}
					m[original].count = m[original].count + 1
				}
			}
		}
	}

	if len(m) == 0 {
		return nil
	}

	result := make([]potentialDoclink, 0, len(m))
	for _, v := range m {
		result = append(result, *v)
	}
	slices.SortFunc(result, func(a, b potentialDoclink) int {
		return strings.Compare(a.original, b.original)
	})
	return result
}
