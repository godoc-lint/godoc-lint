package strict

// (BG: bad godoc)

// Here a number of combinations of godocs and deprecations are tested.

// Deprecated: with whitespace
//
// godoc
const SingleSingleFooBG = 0

const (
	// Deprecated:with no whitespace
	//
	// godoc
	MultiSingleFooBG = 0
)

// Deprecated: with whitespace
//
// Deprecated:with no whitespace
//
// godoc
const SingleSingleFooVarBG = 0

var (
	// Deprecated: with whitespace
	//
	// godoc
	//
	// Deprecated:with no whitespace
	MultiSingleFooVarBG = 0
)

// Deprecated:with no whitespace
//
// godoc
//
// Deprecated: with whitespace
type TSingleFooBG int

type (
	// godoc
	//
	// Deprecated: with whitespace
	TMultiFooBG int

	// godoc
	//
	// Deprecated:with no whitespace
	TMultiBarBG int
)

// Deprecated: with whitespace
//
// Deprecated:with no whitespace
func FooFuncBG() {}

// Deprecated: with whitespace
//
// Deprecated:with no whitespace
//
// Deprecated: with whitespace
//
// godoc
type TFooBG int

// godoc
//
// Deprecated: with whitespace
//
// Deprecated:with no whitespace
//
// Deprecated: with whitespace
func (*TFooBG) FooFuncBG() {}
