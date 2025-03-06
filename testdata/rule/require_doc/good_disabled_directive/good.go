package good

// (NG: no godoc)

//godoclint:disable
const SingleSingleFooDisabledNG = 0

//godoclint:disable require-doc
const SingleMultiFooDisabledNG, SingleMultiBarDisabledNG = 0, 0

const (
	//godoclint:disable
	MultiSingleFooDisabledNG = 0
)

const (
	//godoclint:disable require-doc
	MultiMultiFooDisabledNG, MultiMultiBarDisabledNG = 0, 0
)

//godoclint:disable
type SingleTFooDisabledNG int

type (
	//godoclint:disable
	MultiTFooDisabledNG int
)

//godoclint:disable
func FooDisabledNG() {}

//godoclint:disable
type TFooDisabledNG string

//godoclint:disable require-doc
func (*TFooDisabledNG) TFooBarNG() {}

//godoclint:disable
const singleSingleFooDisabledNG = 0

//godoclint:disable require-doc
const singleMultiFooDisabledNG, singleMultiBarDisabledNG = 0, 0

const (
	//godoclint:disable
	multiSingleFooDisabledNG = 0
)

const (
	//godoclint:disable require-doc
	multiMultiFooDisabledNG, multiMultiBarDisabledNG = 0, 0
)

//godoclint:disable
type singleTFooDisabledNG int

type (
	//godoclint:disable
	multiTFooDisabledNG int
)

//godoclint:disable
func fooDisabledNG() {}

//godoclint:disable
type tFooDisabledNG string

//godoclint:disable require-doc
func (*tFooDisabledNG) tFooBarNG() {}
