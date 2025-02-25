package decl_type

// (NG: No Godoc immediately before declaration)
// (Here NG symbols are defined as type aliases)

// some godoc
type SingleFoo int

type SingleFooNG = int

// parent godoc
type (
	// some godoc
	SingleFooMultiline int
)

// parent godoc
type (
	SingleFooMultilineNG = int
)

// parent godoc
type (
	// first godoc
	MultiFooMultiline int
	// second godoc
	MultiFooMultiline2   int
	MultiFooMultiline3NG = int
)

// parent godoc
type (
	MultiFooMultilineNG  = int
	MultiFooMultiline2NG = int
)
