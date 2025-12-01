package good

// godoc
const FooTest = 0

// godoc
type TFooTest int

// godoc
func FFooTest() {}

// godoc
func (*TFooTest) FooFooTest() {}

const fooTest = 0

type tFooTest int

func fFooTest() {}

func (*tFooTest) fooFooTest() {}

// Below blank declaration has no godoc, but it should be ignored.

var _ = 0
