package model

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

// Inspector defines a pre-run inspector.
type Inspector interface {
	// GetAnalyzer returns the underlying analyzer.
	GetAnalyzer() *analysis.Analyzer
}

// InspectorResult represents the result of the inspector analysis.
type InspectorResult struct {
	// Files provides extracted information per AST file.
	Files map[*ast.File]*FileInspection
}

type FileInspection struct {
	// DisabledRulesMap contains information about rules disabled at top level.
	DisabledRules InspectorResultDisableRules

	// SymbolDecl represents symbols declared in the package file.
	SymbolDecl []SymbolDecl
}

// InspectorResultDisableRules contains the list of disabled rules.
type InspectorResultDisableRules struct {
	// All indicates whether all rules are disabled.
	All bool

	// Rules is the set of rules disabled.
	//
	// If all rules are disable, this will be nil.
	Rules map[string]struct{}
}

// SymbolDeclKind is the enum type for the symbol declarations.
type SymbolDeclKind string

const (
	SymbolDeclKindBad   SymbolDeclKind = "bad"
	SymbolDeclKindFunc  SymbolDeclKind = "func"
	SymbolDeclKindConst SymbolDeclKind = "const"
	SymbolDeclKindType  SymbolDeclKind = "type"
	SymbolDeclKindVar   SymbolDeclKind = "var"
)

// SymbolDecl represents a top level declaration.
type SymbolDecl struct {
	// Decl is the underlying declaration node.
	Decl ast.Decl

	// Kind is the declaration kind (e.g., func or type).
	Kind SymbolDeclKind

	// IsTypeAlias indicates that the type symbol is an alias. For example:
	//
	//  type Foo = int
	//
	// This is always false for non-type declaration (e.g., const or var).
	IsTypeAlias bool

	// Name is the name of the declared symbol.
	Name string

	// MultiNameDecl determines whether the symbol is declared as part of a
	// multi-name declaration spec; For example:
	//
	//   const foo, bar = 0, 0
	//
	// This field is only valid for const, var, or type declarations.
	MultiNameDecl bool

	// Doc is the comment group associated to the symbol.
	Doc *ast.CommentGroup

	// Doc is the comment group associated to the parent declaration. For
	// instance:
	//
	//  // parent godoc
	//  const (
	//      // godoc
	//      Foo = 0
	//  )
	ParentDoc *ast.CommentGroup

	// DisabledRules is the set of rules disabled via inline directives.
	DisabledRules InspectorResultDisableRules
}
