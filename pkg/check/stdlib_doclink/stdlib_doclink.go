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
	enforceRepeats := actx.Config.GetRuleOptions().StdlibDoclinkEnforceRepeats

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
		checkStdlibDoclink(actx, doc, enforceRepeats)
	}
	return nil
}

func walkBlocks(blocks []gdc.Block, fn func(node any) bool) {
	var walk func(node any)
	walk = func(node any) {
		if !fn(node) {
			return
		}

		switch n := node.(type) {
		case *gdc.Code:
			// Do nothing.
		case *gdc.Heading:
			for _, t := range n.Text {
				walk(t)
			}
		case *gdc.List:
			for _, i := range n.Items {
				walk(i)
			}
		case *gdc.ListItem:
			for _, i := range n.Content {
				walk(i)
			}
		case *gdc.Paragraph:
			for _, t := range n.Text {
				walk(t)
			}
		case *gdc.Link:
			for _, t := range n.Text {
				walk(t)
			}
		case *gdc.DocLink:
			for _, t := range n.Text {
				walk(t)
			}
		}
	}

	for _, b := range blocks {
		walk(b)
	}
}

func checkStdlibDoclink(actx *model.AnalysisContext, doc *model.CommentGroup, enforceRepeats bool) {
	if doc.DisabledRules.All || doc.DisabledRules.Rules.Has(StdlibDoclinkRule) {
		return
	}

	var doclinks map[string]struct{}
	if !enforceRepeats {
		doclinks = make(map[string]struct{}, 5)
		walkBlocks(doc.Parsed.Content, func(node any) bool {
			dc, ok := node.(*gdc.DocLink)
			if !ok {
				return true
			}

			if dc.Recv != "" && dc.Name != "" {
				// pkg.name.name
				k := fmt.Sprintf("%s.%s.%s", dc.ImportPath, dc.Recv, dc.Name)
				doclinks[k] = struct{}{}
			} else if dc.Name != "" {
				// pkg.name
				k := fmt.Sprintf("%s.%s", dc.ImportPath, dc.Name)
				doclinks[k] = struct{}{}
			}
			return false
		})
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

	pds := findPotentialDoclinks(text)
	if len(pds) == 0 {
		return
	}

	if !enforceRepeats {
		refined := make([]potentialDoclink, 0, len(pds))
		for _, pd := range pds {
			if _, ok := doclinks[pd.originalNoStar]; ok {
				continue
			}
			refined = append(refined, pd)
		}
		pds = refined
	}

	for _, pd := range pds {
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
	original       string
	originalNoStar string
	count          int
	doclink        string
	kind           internal.SymbolKind
}

var potentialDoclinkRE = sync.OnceValue(func() *regexp.Regexp {
	quotedPkgPaths := make([]string, 0, len(stdlib()))
	for _, s := range stdlib() {
		quotedPkgPaths = append(quotedPkgPaths, regexp.QuoteMeta(s.Path))
	}

	// The package-only case is not a match due to potential false positives matching
	// common words like "bytes", or "time". Even cases like "net/http" can be used
	// in a legitimate godoc text.
	//
	// The error is never non-nil due to tests.
	stdlibPkgRE, _ := regexp.Compile(fmt.Sprintf(`(?m)(?:^|[^\n[*])(\*?)\b(\*?)(%s)\.([a-zA-Z0-9_]+)(?:\.([a-zA-Z0-9_]+))?\b`, strings.Join(quotedPkgPaths, "|")))

	return stdlibPkgRE
})

func findPotentialDoclinks(text string) []potentialDoclink {
	m := make(map[string]*potentialDoclink, 5)

	matches := potentialDoclinkRE().FindAllStringSubmatch(text, -1)
	for _, match := range matches {
		star := match[1]
		pkg := match[2]
		name1 := match[3]
		name2 := match[4]

		if pkg != "" && name1 != "" && name2 != "" {
			// pkg.name1.name2
			if s, ok := stdlib()[pkg]; ok {
				if kind, ok := s.Symbols[name1+"."+name2]; ok {
					originalNoStar := fmt.Sprintf("%s.%s.%s", pkg, name1, name2)
					original := star + originalNoStar
					if _, ok := m[originalNoStar]; !ok {
						doclink := fmt.Sprintf("[%s]", original)
						m[original] = &potentialDoclink{original: original, originalNoStar: originalNoStar, doclink: doclink, kind: kind}
					}
					m[original].count = m[original].count + 1
				}
			}
		} else if pkg != "" && name1 != "" && name2 == "" {
			// pkg.name1
			if s, ok := stdlib()[pkg]; ok {
				if kind, ok := s.Symbols[name1]; ok {
					originalNoStar := fmt.Sprintf("%s.%s", pkg, name1)
					original := star + originalNoStar
					if _, ok := m[originalNoStar]; !ok {
						doclink := fmt.Sprintf("[%s]", original)
						m[original] = &potentialDoclink{original: original, originalNoStar: originalNoStar, doclink: doclink, kind: kind}
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
