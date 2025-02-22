package inspect

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
	"regexp"
	"slices"
	"strings"

	"golang.org/x/tools/go/analysis"

	"github.com/godoc-lint/godoc-lint/pkg/model"
	"github.com/godoc-lint/godoc-lint/pkg/util"
)

const (
	metaName = "godoclint_inspect"
	metaDoc  = "Pre-run inspector for godoclint."
	metaURL  = "https://github.com/godoc-lint/godoc-lint"
)

// Inspector implements the godoc-lint pre-run inspector.
type Inspector struct {
	cb model.ConfigBuilder

	analyzer *analysis.Analyzer
}

// NewInspector returns a new instance of the inspector.
func NewInspector(cb model.ConfigBuilder) *Inspector {
	result := &Inspector{
		cb: cb,
		analyzer: &analysis.Analyzer{
			Name:       metaName,
			Doc:        metaDoc,
			URL:        metaURL,
			ResultType: reflect.TypeOf(new(model.InspectorResult)),
		},
	}
	result.analyzer.Run = result.run
	return result
}

// GetAnalyzer returns the underlying analyzer.
func (i *Inspector) GetAnalyzer() *analysis.Analyzer {
	return i.analyzer
}

var topLevelOrphanCommentGroupPattern = regexp.MustCompile(`(?m)(?:^//.*\r?\n)+(?:\r?\n|\z)`)
var disableDirectivePattern = regexp.MustCompile(`(?m)//godoclint:disable(?: *(.+))?$`)

func (i *Inspector) run(pass *analysis.Pass) (any, error) {
	cfg := i.cb.MustGetConfig()

	inspect := func(f *ast.File) (*model.FileInspection, error) {
		ft := util.GetPassFileToken(f, pass)
		if ft == nil {
			return nil, nil
		}

		raw, err := pass.ReadFile(ft.Name())
		if err != nil {
			return nil, fmt.Errorf("cannot read file %q: %v", ft.Name(), err)
		}

		// Extract top-level //godoclint:disable directives.
		disabledRules := model.InspectorResultDisableRules{}
		for _, match := range topLevelOrphanCommentGroupPattern.FindAll(raw, -1) {
			d := extractDisableDirectivesInComment(string(match))
			if d.All {
				disabledRules.All = true
			}

			if len(d.Rules) > 0 {
				if disabledRules.Rules == nil {
					disabledRules.Rules = make(map[string]struct{}, len(d.Rules))
				}
				for r := range d.Rules {
					disabledRules.Rules[r] = struct{}{}
				}
			}
		}

		// Extract top-level symbol declarations.
		decls := make([]model.SymbolDecl, 0, len(f.Decls))
		for _, d := range f.Decls {
			switch dt := d.(type) {
			case *ast.FuncDecl:
				decls = append(decls, model.SymbolDecl{
					Decl: d,
					Kind: model.SymbolDeclKindFunc,
					Name: dt.Name.Name,
					Doc:  dt.Doc,
				})
			case *ast.BadDecl:
				decls = append(decls, model.SymbolDecl{
					Decl: d,
					Kind: model.SymbolDeclKindBad,
				})
			case *ast.GenDecl:
				switch dt.Tok {
				case token.CONST, token.VAR:
					kind := model.SymbolDeclKindConst
					if dt.Tok == token.VAR {
						kind = model.SymbolDeclKindVar
					}
					if dt.Lparen == token.NoPos {
						// cases:
						// const ... (single line)
						// var ... (single line)

						spec := dt.Specs[0].(*ast.ValueSpec)
						if len(spec.Names) == 1 {
							// cases:
							// const foo = 0
							// var foo = 0
							decls = append(decls, model.SymbolDecl{
								Decl: d,
								Kind: kind,
								Name: spec.Names[0].Name,
								Doc:  dt.Doc,
							})
						} else {
							// cases:
							// const foo, bar = 0, 0
							// var foo, bar = 0, 0
							for _, n := range spec.Names {
								decls = append(decls, model.SymbolDecl{
									Decl:          d,
									Kind:          kind,
									Name:          n.Name,
									Doc:           dt.Doc,
									MultiNameDecl: true,
								})
							}
						}
					} else {
						// cases:
						// const (
						//     foo = 0
						// )
						// var (
						//     foo = 0
						// )
						// const (
						//     foo, bar = 0, 0
						// )
						// var (
						//     foo, bar = 0, 0
						// )

						for _, s := range dt.Specs {
							spec := s.(*ast.ValueSpec)
							for _, n := range spec.Names {
								decls = append(decls, model.SymbolDecl{
									Decl:          d,
									Kind:          kind,
									Name:          n.Name,
									Doc:           spec.Doc,
									ParentDoc:     dt.Doc,
									MultiNameDecl: len(spec.Names) > 1,
								})
							}
						}
					}
				case token.TYPE:
					if dt.Lparen == token.NoPos {
						// case:
						// type foo int

						spec := dt.Specs[0].(*ast.TypeSpec)
						decls = append(decls, model.SymbolDecl{
							Decl:        d,
							Kind:        model.SymbolDeclKindType,
							IsTypeAlias: spec.Assign != token.NoPos,
							Name:        spec.Name.Name,
							Doc:         dt.Doc,
						})
					} else {
						// case:
						// type (
						//     foo int
						// )

						for _, s := range dt.Specs {
							spec := s.(*ast.TypeSpec)
							decls = append(decls, model.SymbolDecl{
								Decl:        d,
								Kind:        model.SymbolDeclKindType,
								IsTypeAlias: spec.Assign != token.NoPos,
								Name:        spec.Name.Name,
								Doc:         spec.Doc,
								ParentDoc:   dt.Doc,
							})
						}
					}
				default:
					continue
				}
			}
		}

		// Extract declaration-scope disable directives.
		for ix, decl := range decls {
			var docs []*ast.Comment

			if decl.Doc != nil {
				docs = slices.Concat(docs, decl.Doc.List)
			}
			if decl.ParentDoc != nil {
				docs = slices.Concat(docs, decl.ParentDoc.List)
			}
			lines := make([]string, 0, len(docs))
			for _, l := range docs {
				lines = append(lines, l.Text)
			}
			text := strings.Join(lines, "\n")
			decls[ix].DisabledRules = extractDisableDirectivesInComment(text)
		}

		return &model.FileInspection{
			DisabledRules: disabledRules,
			SymbolDecl:    decls,
		}, nil
	}

	result := &model.InspectorResult{
		Files: make(map[*ast.File]model.FileInspection, len(pass.Files)),
	}

	for _, f := range pass.Files {
		if f.Pos() == token.NoPos {
			continue
		}
		ft := pass.Fset.File(f.Pos())
		if ft == nil {
			continue
		}
		if !cfg.IsPathApplicable(ft.Name()) {
			continue
		}

		if fi, err := inspect(f); err != nil {
			return nil, fmt.Errorf("inspector failed: %w", err)
		} else {
			result.Files[f] = *fi
		}
	}
	return result, nil
}

func extractDisableDirectivesInComment(s string) model.InspectorResultDisableRules {
	result := model.InspectorResultDisableRules{}
	for _, directive := range disableDirectivePattern.FindAllStringSubmatch(s, -1) {
		args := string(directive[1])
		if args == "" {
			result.All = true
			continue
		}
		ruleNames := strings.Split(strings.TrimSpace(args), " ")
		if result.Rules == nil {
			result.Rules = make(map[string]struct{}, len(ruleNames))
		}
		for _, r := range ruleNames {
			result.Rules[r] = struct{}{}
		}
	}
	return result
}
