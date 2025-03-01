package decl_type

// (NG: No Godoc immediately before declaration)
// (Here NG symbols are defined as type aliases)

// some godoc
type SingleFoo int // trailing doc

type SingleFooNG = int

// parent godoc
type (
	// some godoc
	SingleFooMultiline int // trailing doc
)

// parent godoc
type (
	SingleFooMultilineNG = int
)

// parent godoc
type (
	// first godoc
	MultiFooMultiline int // trailing doc
	// second godoc
	MultiFooMultiline2   int // trailing doc
	MultiFooMultiline3NG = int
)

// parent godoc
type (
	MultiFooMultilineNG  = int
	MultiFooMultiline2NG = int
)
