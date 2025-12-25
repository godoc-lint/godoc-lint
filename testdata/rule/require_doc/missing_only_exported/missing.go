package missing

// (NG: no godoc)

// The //foo:bar directives mark the trailing comment as a directive so they're
// not parsed as a normal trailing comment group.

const SingleSingleFooNG = 0 //foo:bar // want `symbol should have a godoc \("SingleSingleFooNG"\)`

const SingleMultiFooNG, SingleMultiBarNG = 0, 0 //foo:bar // want `symbol should have a godoc \("SingleMultiFooNG"\)` `symbol should have a godoc \("SingleMultiBarNG"\)`

const (
	MultiSingleFooNG = 0 //foo:bar // want `symbol should have a godoc \("MultiSingleFooNG"\)`
)

const (
	MultiMultiFooNG, MultiMultiBarNG = 0, 0 //foo:bar // want `symbol should have a godoc \("MultiMultiFooNG"\)` `symbol should have a godoc \("MultiMultiBarNG"\)`
)

type SingleTFooNG int //foo:bar // want `symbol should have a godoc \("SingleTFooNG"\)`

type (
	MultiTFooNG int //foo:bar // want `symbol should have a godoc \("MultiTFooNG"\)`
)

func FooNG() {} //foo:bar // want `symbol should have a godoc \("FooNG"\)`

type TFooNG string //foo:bar // want `symbol should have a godoc \("TFooNG"\)`

func (*TFooNG) TFooBarNG() {} //foo:bar // want `symbol should have a godoc \("TFooBarNG"\)`

func (*TFooNG) tFooBarNG() {}

const singleSingleFooNG = 0

const singleMultiFooNG, singleMultiBarNG = 0, 0

const (
	multiSingleFooNG = 0
)

const (
	multiMultiFooNG, multiMultiBarNG = 0, 0
)

type singleTFooNG int

type (
	multiTFooNG int
)

func funcFooNG() {}

type tFooNG string

func (*tFooNG) tFooBarNG() {}

//foo:bar // unexported receiver
func (*tFooNG) TFooBarNG() {}
