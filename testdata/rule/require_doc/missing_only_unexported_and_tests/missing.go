package missing

// (NG: no godoc)

// The //foo:bar directives mark the trailing comment as a directive so they're
// not parsed as a normal trailing comment group.

const SingleSingleFooNG = 0

const SingleMultiFooNG, SingleMultiBarNG = 0, 0

const (
	MultiSingleFooNG = 0
)

const (
	MultiMultiFooNG, MultiMultiBarNG = 0, 0
)

type SingleTFooNG int

type (
	MultiTFooNG int
)

func FooNG() {}

type TFooNG string

func (*TFooNG) TFooBarNG() {}

func (*TFooNG) tFooBarNG() {} //foo:bar // want `symbol should have a godoc \("tFooBarNG"\)`

const singleSingleFooNG = 0 //foo:bar // want `symbol should have a godoc \("singleSingleFooNG"\)`

const singleMultiFooNG, singleMultiBarNG = 0, 0 //foo:bar // want `symbol should have a godoc \("singleMultiFooNG"\)` `symbol should have a godoc \("singleMultiBarNG"\)`

const (
	multiSingleFooNG = 0 //foo:bar // want `symbol should have a godoc \("multiSingleFooNG"\)`
)

const (
	multiMultiFooNG, multiMultiBarNG = 0, 0 //foo:bar // want `symbol should have a godoc \("multiMultiFooNG"\)` `symbol should have a godoc \("multiMultiBarNG"\)`
)

type singleTFooNG int //foo:bar // want `symbol should have a godoc \("singleTFooNG"\)`

type (
	multiTFooNG int //foo:bar // want `symbol should have a godoc \("multiTFooNG"\)`
)

func funcFooNG() {} //foo:bar // want `symbol should have a godoc \("funcFooNG"\)`

type tFooNG string //foo:bar // want `symbol should have a godoc \("tFooNG"\)`

func (*tFooNG) tFooBarNG() {} //foo:bar // want `symbol should have a godoc \("tFooBarNG"\)`

//foo:bar // unexported receiver
func (*tFooNG) TFooBarNG() {} //foo:bar // want `symbol should have a godoc \("TFooBarNG"\)`
