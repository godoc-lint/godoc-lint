package good_test

// (NG: no godoc)

const FooTestNG = 0

type TFooTestNG int

func FFooTestNG() {}

func (*TFooTestNG) FooFooTest() {}

const fooTestNG = 0

type ttFooTestNG int

func fFooTestNG() {}

func (*ttFooTestNG) fooFooTestNG() {}

// Below blank declaration has no godoc, but it should be ignored.

var _ = 0
