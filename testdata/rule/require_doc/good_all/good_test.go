package good

// (NG: no godoc)

// godoc
const FooTestNG = 0

// godoc
type TFooTestNG int

// godoc
func FFooTestNG() {}

// godoc
func (*TFooTestNG) FooFooTest() {}

// godoc
const fooTestNG = 0

// godoc
type ttFooTestNG int

// godoc
func fFooTestNG() {}

// godoc
func (*ttFooTestNG) fooFooTestNG() {}

// Below blank declaration has no godoc, but it should be ignored.

var _ = 0
