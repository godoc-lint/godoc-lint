package normal

// (BG: bad godoc)

// godoc // want `godoc should start with symbol name \("SingleSingleFooBG"\)`
const SingleSingleFooBG = 0

const (
	// godoc // want `godoc should start with symbol name \("MultiSingleFooBG"\)`
	MultiSingleFooBG = 0
)

// godoc // want `godoc should start with symbol name \("SingleSingleFooVarBG"\)`
const SingleSingleFooVarBG = 0

var (
	// godoc // want `godoc should start with symbol name \("MultiSingleFooVarBG"\)`
	MultiSingleFooVarBG = 0
)

// godoc // want `godoc should start with symbol name \("TSingleFooBG"\)`
type TSingleFooBG int

type (
	// godoc // want `godoc should start with symbol name \("TMultiFooBG"\)`
	TMultiFooBG int

	// godoc // want `godoc should start with symbol name \("TMultiBarBG"\)`
	TMultiBarBG int
)

// godoc // want `godoc should start with symbol name \("FooFuncBG"\)`
func FooFuncBG() {}

// godoc // want `godoc should start with symbol name \("TFooBG"\)`
type TFooBG int

// godoc // want `godoc should start with symbol name \("FooFuncBG"\)`
func (*TFooBG) FooFuncBG() {}
