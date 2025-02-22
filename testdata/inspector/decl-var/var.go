package decl_var

// (NG: No Godoc immediately before declaration)

// some godoc
var SingleSingleFoo = 0

var SingleSingleFooNG = 0

// some godoc
var SingleDoubleFoo, SingleDoubleBar = 0, 0

var SingleDoubleFooNG, SingleDoubleBarNG = 0, 0

// parent godoc
var (
	// some godoc
	SingleSingleFooMultiline = 0
)

// parent godoc
var (
	SingleSingleFooMultilineNG = 0
)

// parent godoc
var (
	// first godoc
	MultiSingleFooMultiline = 0
	// second godoc
	MultiSingleFooMultiline2   = 0
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
	SingleDoubleFooMultiline, SingleDoubleBarMultiline = 0, 0
)

// parent godoc
var (
	SingleDoubleFooMultilineNG, SingleDoubleBarMultilineNG = 0, 0
)

// parent godoc
var (
	// first godoc
	MultiDoubleFooMultiline, MultiDoubleBarMultiline = 0, 0
	// second godoc
	MultiDoubleFooMultiline2, MultiDoubleBarMultiline2     = 0, 0
	MultiDoubleFooMultiline3NG, MultiDoubleBarMultiline3NG = 0, 0
)

// parent godoc
var (
	MultiDoubleFooMultilineNG, MultiDoubleBarMultilineNG   = 0, 0
	MultiDoubleFooMultiline2NG, MultiDoubleBarMultiline2NG = 0, 0
)
