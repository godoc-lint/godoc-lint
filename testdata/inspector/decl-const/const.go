package decl_const

// (NG: No Godoc immediately before declaration)

// some godoc
const SingleSingleFoo = 0

const SingleSingleFooNG = 0

// some godoc
const SingleDoubleFoo, SingleDoubleBar = 0, 0

const SingleDoubleFooNG, SingleDoubleBarNG = 0, 0

// parent godoc
const (
	// some godoc
	SingleSingleFooMultiline = 0
)

// parent godoc
const (
	SingleSingleFooMultilineNG = 0
)

// parent godoc
const (
	// first godoc
	MultiSingleFooMultiline = 0
	// second godoc
	MultiSingleFooMultiline2   = 0
	MultiSingleFooMultiline3NG = 0
)

// parent godoc
const (
	MultiSingleFooMultilineNG  = 0
	MultiSingleFooMultiline2NG = 0
)

// parent godoc
const (
	// some godoc
	SingleDoubleFooMultiline, SingleDoubleBarMultiline = 0, 0
)

// parent godoc
const (
	SingleDoubleFooMultilineNG, SingleDoubleBarMultilineNG = 0, 0
)

// parent godoc
const (
	// first godoc
	MultiDoubleFooMultiline, MultiDoubleBarMultiline = 0, 0
	// second godoc
	MultiDoubleFooMultiline2, MultiDoubleBarMultiline2     = 0, 0
	MultiDoubleFooMultiline3NG, MultiDoubleBarMultiline3NG = 0, 0
)

// parent godoc
const (
	MultiDoubleFooMultilineNG, MultiDoubleBarMultilineNG   = 0, 0
	MultiDoubleFooMultiline2NG, MultiDoubleBarMultiline2NG = 0, 0
)
