package good

// (NG: no godoc)

//godoclint:disable require-doc

const FooTestFileNG = 0

type TFooTestFileNG int

func FFooTestFileNG() {}

func (*TFooTestFileNG) FooFooTest() {}

func (*TFooTestFileNG) fooFooTest() {}

const fooTestFileNG = 0

type ttFooTestFileNG int

func fFooTestFileNG() {}

func (*ttFooTestFileNG) fooFooTestFileNG() {}

func (*ttFooTestFileNG) FooFooTestFileNG() {}
