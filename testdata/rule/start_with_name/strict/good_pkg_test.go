package strict_test

// (BG: bad godoc)

// godoc
const TestSingleSingleFooBG = 0

const (
	// godoc
	TestMultiSingleFooBG = 0
)

// godoc
const TestSingleSingleFooVarBG = 0

var (
	// godoc
	TestMultiSingleFooVarBG = 0
)

// godoc
type TestTSingleFooBG int

type (
	// godoc
	TestTMultiFooBG int

	// godoc
	TestTMultiBarBG int
)

// godoc
func TESTFooFuncBG() {}

// godoc
type TestTFooBG int

// godoc
func (*TestTFooBG) TestFooFuncBG() {}
