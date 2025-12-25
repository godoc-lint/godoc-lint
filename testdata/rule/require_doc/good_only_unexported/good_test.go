package good

const FooTest = 0

type TFooTest int

func FFooTest() {}

func (*TFooTest) FooFooTest() {}

func (*TFooTest) fooFooTest() {}

const fooTest = 0

type tFooTest int

func fFooTest() {}

func (*tFooTest) fooFooTest() {}

func (*tFooTest) FooFooTest() {}

// Below blank declaration has no godoc, but it should be ignored.

var _ = 0
