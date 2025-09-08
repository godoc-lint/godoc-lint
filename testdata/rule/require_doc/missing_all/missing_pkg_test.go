package missing_test

const FooTest = 0 //foo:bar // want `symbol should have a godoc \("FooTest"\)`

type TFooTest int //foo:bar // want `symbol should have a godoc \("TFooTest"\)`

func FFooTest() {} //foo:bar // want `symbol should have a godoc \("FFooTest"\)`

func (*TFooTest) FooFooTest() {} //foo:bar // want `symbol should have a godoc \("FooFooTest"\)`

const fooTest = 0 //foo:bar // want `symbol should have a godoc \("fooTest"\)`

type tFooTest int //foo:bar // want `symbol should have a godoc \("tFooTest"\)`

func fFooTest() {} //foo:bar // want `symbol should have a godoc \("fFooTest"\)`

func (*tFooTest) fooFooTest() {} //foo:bar // want `symbol should have a godoc \("fooFooTest"\)`
