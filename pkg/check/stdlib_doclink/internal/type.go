package internal

import (
	_ "embed"
)

type SymbolKind string

const (
	SymbolKindNA     SymbolKind = ""
	SymbolKindConst  SymbolKind = "c"
	SymbolKindVar    SymbolKind = "v"
	SymbolKindFunc   SymbolKind = "f"
	SymbolKindMethod SymbolKind = "m"
	SymbolKindType   SymbolKind = "t"
)

type StdlibPackage struct {
	Path    string                `json:"path"`
	Name    string                `json:"name"`
	Symbols map[string]SymbolKind `json:"symbols"`
}

type Stdlib map[string]StdlibPackage
