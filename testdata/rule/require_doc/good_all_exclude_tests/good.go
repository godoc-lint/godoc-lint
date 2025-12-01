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

// Below blank declaration has no godoc, but it should be ignored.

var _ = 0
