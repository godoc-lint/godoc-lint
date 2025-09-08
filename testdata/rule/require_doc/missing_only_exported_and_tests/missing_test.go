package missing

const FooTest = 0 //foo:bar // want `symbol should have a godoc \("FooTest"\)`

type TFooTest int //foo:bar // want `symbol should have a godoc \("TFooTest"\)`

func FFooTest() {} //foo:bar // want `symbol should have a godoc \("FFooTest"\)`

func (*TFooTest) FooFooTest() {} //foo:bar // want `symbol should have a godoc \("FooFooTest"\)`

const fooTest = 0

type tFooTest int

func fFooTest() {}

func (*tFooTest) fooFooTest() {}
