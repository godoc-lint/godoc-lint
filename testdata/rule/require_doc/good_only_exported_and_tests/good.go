package good

// (TC: with trailing comment)
// (GD: with godoc)
// (PGD: with parent godoc)
// (All: TC + GD + PGD)

// godoc
const SingleSingleFooGD = 0

// godoc
const SingleMultiFooGD, SingleMultiBarGD = 0, 0

const (
	// godoc
	MultiSingleFooGD = 0
)

const (
	// godoc
	MultiMultiFooGD, MultiMultiBarGD = 0, 0
)

// godoc
type SingleTFooGD int

type (
	// godoc
	MultiTFooGD int
)

const SingleSingleFooTC = 0 // godoc

const SingleMultiFooTC, SingleMultiBarTC = 0, 0 // godoc

const (
	MultiSingleFooTC = 0 // godoc
)

const (
	MultiMultiFooTC, MultiMultiBarTC = 0, 0 // godoc
)

type SingleTFooTC int // godoc

type (
	MultiTFooTC int // godoc
)

// godoc
const (
	MultiSingleFooPGD = 0
)

// godoc
const (
	MultiMultiFooPGD, MultiMultiBarPGD = 0, 0
)

// godoc
type (
	MultiTFooPGD int
)

// godoc
const (
	// godoc
	MultiSingleFooAll = 0 // godoc
)

// godoc
const (
	// godoc
	MultiMultiFooAll, MultiMultiBarAll = 0, 0 // godoc
)

// godoc
type (
	// godoc
	MultiTFooAll int // godoc
)

// godoc
func FuncFoo() {}

// godoc
type TFoo string

// godoc
func (*TFoo) TFooBar() {}

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

// Below blank declaration has no godoc, but it should be ignored.

var _ = 0
