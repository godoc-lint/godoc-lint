package good

// (NG: no godoc)

const FooTestNG = 0

type TFooTestNG int

func FFooTestNG() {}

func (*TFooTestNG) FooFooTest() {}

func (*TFooTestNG) fooFooTest() {}

const fooTestNG = 0

type ttFooTestNG int

func fFooTestNG() {}

func (*ttFooTestNG) fooFooTestNG() {}

func (*ttFooTestNG) FooFooTestNG() {}

// Below blank declaration has no godoc, but it should be ignored.

var _ = 0
