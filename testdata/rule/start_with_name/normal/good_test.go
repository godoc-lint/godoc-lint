package normal

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

// Bad godoc, but should be ignored due to blank identifier.
var _ = 0
