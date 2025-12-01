package normal

// (NG: no godoc)

const SingleSingleFooNG = 0

// SingleSingleFoo has a godoc.
const SingleSingleFoo = 0

const (
	// MultiSingleFoo
	MultiSingleFoo = 0
)

const (
	// This should be fine since it's a multi-name declaration.
	MultiMultiFoo, MultiMultiBar = 0, 0
)

// SingleSingleFooVar has a godoc.
const SingleSingleFooVar = 0

var (
	// MultiSingleFooVar
	MultiSingleFooVar = 0
)

var (
	// This should be fine since it's a multi-name declaration.
	MultiMultiFooVar, MultiMultiBarVar = 0, 0
)

// TSingleFoo has a godoc
type TSingleFoo int

type (
	// TMultiFoo
	TMultiFoo int

	// TMultiBar
	TMultiBar int
)

// FooFunc has a godoc.
func FooFunc() {}

// TFoo
type TFoo int

// FooFunc has a godoc.
func (*TFoo) FooFunc() {}

// Bad godoc, but should be ignored due to blank identifier.
var _ = 0
