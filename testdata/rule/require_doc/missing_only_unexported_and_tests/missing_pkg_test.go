package missing_test

const FooTest = 0

type TFooTest int

func FFooTest() {}

func (*TFooTest) FooFooTest() {}

func (*TFooTest) fooFooTest() {} //foo:bar // want `symbol should have a godoc \("fooFooTest"\)`

const fooTest = 0 //foo:bar // want `symbol should have a godoc \("fooTest"\)`

type tFooTest int //foo:bar // want `symbol should have a godoc \("tFooTest"\)`

func fFooTest() {} //foo:bar // want `symbol should have a godoc \("fFooTest"\)`

func (*tFooTest) fooFooTest() {} //foo:bar // want `symbol should have a godoc \("fooFooTest"\)`

//foo:bar // unexported receiver
func (*tFooTest) FooFooTest() {} //foo:bar // want `symbol should have a godoc \("FooFooTest"\)`
