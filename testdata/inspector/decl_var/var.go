package decl_var

// (NG: No Godoc immediately before declaration)

// some godoc
var SingleSingleFoo = 0 // trailing doc

var SingleSingleFooNG = 0

// some godoc
var SingleDoubleFoo, SingleDoubleBar = 0, 0 // trailing doc

var SingleDoubleFooNG, SingleDoubleBarNG = 0, 0

// parent godoc
var (
	// some godoc
	SingleSingleFooMultiline = 0 // trailing doc
)

// parent godoc
var (
	SingleSingleFooMultilineNG = 0
)

// parent godoc
var (
	// first godoc
	MultiSingleFooMultiline = 0 // trailing doc
	// second godoc
	MultiSingleFooMultiline2   = 0 // trailing doc
	MultiSingleFooMultiline3NG = 0
)

// parent godoc
var (
	MultiSingleFooMultilineNG  = 0
	MultiSingleFooMultiline2NG = 0
)

// parent godoc
var (
	// some godoc
	SingleDoubleFooMultiline, SingleDoubleBarMultiline = 0, 0 // trailing doc
)

// parent godoc
var (
	SingleDoubleFooMultilineNG, SingleDoubleBarMultilineNG = 0, 0
)

// parent godoc
var (
	// first godoc
	MultiDoubleFooMultiline, MultiDoubleBarMultiline = 0, 0 // trailing doc
	// second godoc
	MultiDoubleFooMultiline2, MultiDoubleBarMultiline2     = 0, 0 // trailing doc
	MultiDoubleFooMultiline3NG, MultiDoubleBarMultiline3NG = 0, 0
)

// parent godoc
var (
	MultiDoubleFooMultilineNG, MultiDoubleBarMultilineNG   = 0, 0
	MultiDoubleFooMultiline2NG, MultiDoubleBarMultiline2NG = 0, 0
)
