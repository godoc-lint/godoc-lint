package strict

// SingleSingleFoo has a godoc.
const SingleSingleFoo = 0

const (
	// MultiSingleFoo
	MultiSingleFoo = 0
)

const (
	// This should be fine since it's a multi-name declaration.
	MultiMultiFoo, MultiMultiBar = 0, 0
)

// TSingleFoo has a godoc
type TSingleFoo int

type (
	// TMultiFoo
	TMultiFoo int

	// TMultiBar
	TMultiBar int
)
