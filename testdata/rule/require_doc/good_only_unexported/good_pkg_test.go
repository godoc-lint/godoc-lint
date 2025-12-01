package good_test

const FooTest = 0

type TFooTest int

func FFooTest() {}

func (*TFooTest) FooFooTest() {}

// godoc
const fooTest = 0

// godoc
type tFooTest int

// godoc
func fFooTest() {}

// godoc
func (*tFooTest) fooFooTest() {}

// Below blank declaration has no godoc, but it should be ignored.

var _ = 0
