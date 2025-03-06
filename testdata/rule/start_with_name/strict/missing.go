package strict

// (BG: bad godoc)

// godoc // want `godoc should start with symbol name \(pattern ""\)`
const SingleSingleFooBG = 0

const (
	// godoc // want `godoc should start with symbol name \(pattern ""\)`
	MultiSingleFooBG = 0
)

// godoc // want `godoc should start with symbol name \(pattern ""\)`
const SingleSingleFooVarBG = 0

var (
	// godoc // want `godoc should start with symbol name \(pattern ""\)`
	MultiSingleFooVarBG = 0
)

// godoc // want `godoc should start with symbol name \(pattern ""\)`
type TSingleFooBG int

type (
	// godoc // want `godoc should start with symbol name \(pattern ""\)`
	TMultiFooBG int

	// godoc // want `godoc should start with symbol name \(pattern ""\)`
	TMultiBarBG int
)

// godoc // want `godoc should start with symbol name \(pattern ""\)`
func FooFuncBG() {}

// godoc // want `godoc should start with symbol name \(pattern ""\)`
type TFooBG int

// godoc // want `godoc should start with symbol name \(pattern ""\)`
func (*TFooBG) FooFuncBG() {}
