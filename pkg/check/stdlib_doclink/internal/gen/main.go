// gen is a command package that extracts information about stdlib's top-level
// declarations and generates a JSON.
//
// This package should be compiled into a binary (using the same or newer Go
// compiler) and then invoked in the root of the Go stdlib source tree. The
// result is a JSON file named "stdlib.json" that contains exported symbols
// defined in the standard library. The JSON file is meant to be checked into
// the repository to be used when applying the "require-stdlib-doclink" rule.
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/ast"
	"go/token"
	"os"
	"strings"
	"sync"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/godoc-lint/godoc-lint/pkg/check/stdlib_doclink/internal"
	"github.com/godoc-lint/godoc-lint/pkg/util"
)

func main() {
	singlechecker.Main(&analyzer)
}

type stdlibFact struct{}

// AFact is to make type implement the [analysis.Fact] interface.
func (f *stdlibFact) AFact() {}

var analyzer = analysis.Analyzer{
	Name:             "_stdlib",
	Doc:              "n/a",
	FactTypes:        []analysis.Fact{&stdlibFact{}},
	RunDespiteErrors: true,
	Run:              run,
}

var (
	mu = sync.Mutex{}
	m  = internal.Stdlib{}
)

func run(pass *analysis.Pass) (any, error) {
	path := pass.Pkg.Path()

	if !strings.HasPrefix(path, "std/") ||
		strings.HasSuffix(path, ".test") ||
		strings.HasSuffix(path, "_test") ||
		strings.Contains(path, "/internal/") ||
		strings.HasSuffix(path, "/internal") {
		return nil, nil
	}

	path, _ = strings.CutPrefix(pass.Pkg.Path(), "std/")

	p := &internal.StdlibPackage{
		Path:    path,
		Name:    pass.Pkg.Name(),
		Symbols: make(map[string]internal.SymbolKind, 200),
	}

	for _, f := range pass.Files {
		ft := util.GetPassFileToken(f, pass)
		if ft == nil {
			return nil, nil
		}

		if strings.HasSuffix(ft.Name(), "_test.go") {
			continue
		}

		raw, err := pass.ReadFile(ft.Name())
		if err != nil {
			return nil, fmt.Errorf("cannot read file %q: %v", ft.Name(), err)
		}

		_ = raw

		for _, d := range f.Decls {
			switch dt := d.(type) {
			case *ast.FuncDecl:
				if !ast.IsExported(dt.Name.Name) {
					continue
				}
				if dt.Recv == nil {
					key := dt.Name.Name
					p.Symbols[key] = internal.SymbolKindFunc
				} else {
					if se, ok := dt.Recv.List[0].Type.(*ast.StarExpr); ok {
						if ident, ok := se.X.(*ast.Ident); ok {
							if ast.IsExported(ident.Name) {
								key := ident.Name + "." + dt.Name.Name
								p.Symbols[key] = internal.SymbolKindMethod
							}
						}
					} else if ident, ok := dt.Recv.List[0].Type.(*ast.Ident); ok {
						if ast.IsExported(ident.Name) {
							key := ident.Name + "." + dt.Name.Name
							p.Symbols[key] = internal.SymbolKindMethod
						}
					}
				}
			case *ast.GenDecl:
				switch dt.Tok {
				case token.CONST, token.VAR:
					kind := internal.SymbolKindConst
					if dt.Tok == token.VAR {
						kind = internal.SymbolKindVar
					}
					for _, spec := range dt.Specs {
						for _, name := range spec.(*ast.ValueSpec).Names {
							if ast.IsExported(name.Name) {
								key := name.Name
								p.Symbols[key] = kind
							}
						}
					}
				case token.TYPE:
					for _, spec := range dt.Specs {
						if ast.IsExported(spec.(*ast.TypeSpec).Name.Name) {
							key := spec.(*ast.TypeSpec).Name.Name
							p.Symbols[key] = internal.SymbolKindType
						}
					}
				}
			}
		}
	}

	if len(p.Symbols) == 0 {
		return nil, nil
	}

	mu.Lock()
	defer mu.Unlock()

	if p0, ok := m[path]; ok {
		if p0.Name != p.Name {
			panic(fmt.Sprintf("inconsistent package name for %q: %q vs %q", path, p0.Name, p.Name))
		}

		for k := range p.Symbols {
			p0.Symbols[k] = p.Symbols[k]
		}
		p = p0
	}

	m[path] = p

	buf := bytes.NewBuffer(nil)
	if err := json.NewEncoder(buf).Encode(m); err != nil {
		return nil, err
	}
	_ = os.WriteFile("stdlib.json", buf.Bytes(), 0o644)

	return nil, nil
}
