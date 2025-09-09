package strict

// (BGI: bad godoc with invalid deprecation)

// Here, combinations of godocs and deprecations are tested on various declaration kinds.

// Deprecated:invalid deprecation // want `godoc should start with symbol name \("SingleSingleFooBGI"\)`
const SingleSingleFooBGI = 0

const (
	// Deprecated:invalid deprecation // want `godoc should start with symbol name \("MultiSingleFooBGI"\)`
	//
	// bad godoc
	MultiSingleFooBGI = 0
)

// bad godoc // want `godoc should start with symbol name \("SingleSingleFooVarBGI"\)`
//
// Deprecated:invalid deprecation
const SingleSingleFooVarBGI = 0

var (
	// Deprecated:invalid deprecation // want `godoc should start with symbol name \("MultiSingleFooVarBGI"\)`
	//
	// bad godoc
	//
	// Deprecated:invalid deprecation
	MultiSingleFooVarBGI = 0
)

// bad godoc // want `godoc should start with symbol name \("TSingleFooBGI"\)`
//
// Deprecated:invalid deprecation
//
// Deprecated:invalid deprecation
type TSingleFooBGI int

type (
	// Deprecated:invalid deprecation // want `godoc should start with symbol name \("TMultiFooBGI"\)`
	//
	// Deprecated:invalid deprecation
	//
	// bad godoc
	TMultiFooBGI int

	// Deprecated:invalid deprecation // want `godoc should start with symbol name \("TMultiBarBGI"\)`
	//
	// Deprecated:invalid deprecation
	//
	// bad godoc
	TMultiBarBGI int
)

// Deprecated:invalid deprecation // want `godoc should start with symbol name \("FooFuncBGI"\)`
//
// bad godoc
//
// Deprecated:valid deprecation
func FooFuncBGI() {}

// This is not a valid deprecation:

// Deprecated // want `godoc should start with symbol name \("TFooBGI"\)`
type TFooBGI int

// This is not a valid deprecation:

// The //foo:bar directives mark the trailing comment as a directive so they're
// not parsed as a normal trailing comment group.

// Deprecated://foo:bar // want `godoc should start with symbol name \("FooFuncBGI"\)`
func (*TFooBGI) FooFuncBGI() {}
