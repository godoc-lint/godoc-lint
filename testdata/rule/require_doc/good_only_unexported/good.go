package good

// (TC: with trailing comment)
// (GD: with godoc)
// (PGD: with parent godoc)
// (All: TC + GD + PGD)

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

func FuncFooNG() {}

type TFooNG string

func (*TFooNG) TFooBarNG() {}

// godoc
func (*TFooNG) tFooBarNG() {}

// godoc
const singleSingleFooGD = 0

// godoc
const singleMultiFooGD, singleMultiBarGD = 0, 0

const (
	// godoc
	multiSingleFooGD = 0
)

const (
	// godoc
	multiMultiFooGD, multiMultiBarGD = 0, 0
)

// godoc
type singleTFooGD int

type (
	// godoc
	multiTFooGD int
)

const singleSingleFooTC = 0 // godoc

const singleMultiFooTC, singleMultiBarTC = 0, 0 // godoc

const (
	multiSingleFooTC = 0 // godoc
)

const (
	multiMultiFooTC, multiMultiBarTC = 0, 0 // godoc
)

type singleTFooTC int // godoc

type (
	multiTFooTC int // godoc
)

// godoc
const (
	multiSingleFooPGD = 0
)

// godoc
const (
	multiMultiFooPGD, multiMultiBarPGD = 0, 0
)

// godoc
type (
	multiTFooPGD int
)

// godoc
const (
	// godoc
	multiSingleFooAll = 0 // godoc
)

// godoc
const (
	// godoc
	multiMultiFooAll, multiMultiBarAll = 0, 0 // godoc
)

// godoc
type (
	// godoc
	multiTFooAll int // godoc
)

// godoc
func funcFoo() {}

// godoc
type tFoo string

// godoc
func (*tFoo) tFooBar() {}

// godoc
func (*tFoo) TFooBar() {}

// Below blank declaration has no godoc, but it should be ignored.

var _ = 0
