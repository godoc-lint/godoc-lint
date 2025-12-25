package good

// (NG: no godoc)

//godoclint:disable
const FooTestNG = 0

//godoclint:disable require-doc
type TFooTestNG int

//godoclint:disable
func FFooTestNG() {}

//godoclint:disable require-doc
func (*TFooTestNG) FooFooTest() {}

//godoclint:disable require-doc
func (*TFooTestNG) fooFooTest() {}

//godoclint:disable
const fooTestNG = 0

//godoclint:disable require-doc
type ttFooTestNG int

//godoclint:disable
func fFooTestNG() {}

//godoclint:disable require-doc
func (*ttFooTestNG) fooFooTestNG() {}

//godoclint:disable require-doc
func (*ttFooTestNG) FooFooTestNG() {}
