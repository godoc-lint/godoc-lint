package strict

// (BG: bad godoc)

// Here, combinations of godocs and deprecations are tested on various declaration kinds.

// Deprecated: valid deprecation
const SingleSingleFooBG = 0

const (
	// Deprecated: valid deprecation
	//
	// bad godoc but okay due to the valid deprecation marker
	MultiSingleFooBG = 0
)

// bad godoc but okay due to the valid deprecation marker
//
// Deprecated: valid deprecation
const SingleSingleFooVarBG = 0

var (
	// Deprecated: valid deprecation
	//
	// bad godoc but okay due to the valid deprecation marker
	//
	// Deprecated: valid deprecation
	MultiSingleFooVarBG = 0
)

// bad godoc but okay due to the valid deprecation marker
//
// Deprecated: valid deprecation
//
// Deprecated: valid deprecation
type TSingleFooBG int

type (
	// Deprecated: valid deprecation
	//
	// Deprecated: valid deprecation
	//
	// bad godoc but okay due to the valid deprecation marker
	TMultiFooBG int

	// Deprecated: valid deprecation
	//
	// Deprecated:invalid deprecation
	//
	// bad godoc but okay due to the valid deprecation marker
	TMultiBarBG int
)

// Deprecated:invalid deprecation
//
// bad godoc but okay due to the valid deprecation marker
//
// Deprecated: valid deprecation
func FooFuncBG() {}

// Deprecated:invalid deprecation
//
// Deprecated:invalid deprecation
//
// bad godoc but okay due to the valid deprecation marker
//
// Deprecated: valid deprecation
//
// Deprecated:invalid deprecation
type TFooBG int

// Deprecated:invalid deprecation
//
// Deprecated:invalid deprecation
//
// bad godoc but okay due to the valid deprecation marker
//
// Deprecated:invalid deprecation
//
// Deprecated: valid deprecation
func (*TFooBG) FooFuncBG() {}
