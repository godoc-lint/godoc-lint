package good_test

const FooTest = 0

type TFooTest int

func FFooTest() {}

func (*TFooTest) FooFooTest() {}

// godoc
func (*TFooTest) fooFooTest() {}

// godoc
const fooTest = 0

// godoc
type tFooTest int

// godoc
func fFooTest() {}

// godoc
func (*tFooTest) fooFooTest() {}

// godoc
//
//foo:bar // unexported receiver
func (*tFooTest) FooFooTest() {}

// Below blank declaration has no godoc, but it should be ignored.

var _ = 0
