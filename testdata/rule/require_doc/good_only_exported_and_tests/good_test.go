package good

// godoc
const FooTest = 0

// godoc
type TFooTest int

// godoc
func FFooTest() {}

// godoc
func (*TFooTest) FooFooTest() {}

func (*TFooTest) fooFooTest() {}

const fooTest = 0

type tFooTest int

func fFooTest() {}

func (*tFooTest) fooFooTest() {}

//foo:bar // unexported receiver
func (*tFooTest) FooFooTest() {}

// Below blank declaration has no godoc, but it should be ignored.

var _ = 0
