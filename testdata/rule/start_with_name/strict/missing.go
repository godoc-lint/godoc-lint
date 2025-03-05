package strict

// (BG: bad godoc)

// godoc // want `godoc should start with symbol name \(pattern ""\)`
const SingleSingleFooBG = 0

const (
	// godoc // want `godoc should start with symbol name \(pattern ""\)`
	MultiSingleFooBG = 0
)

const (
	// This should be fine since it's a multi-name declaration.
	MultiMultiFooBG, MultiMultiBarBG = 0, 0
)

// godoc // want `godoc should start with symbol name \(pattern ""\)`
type TSingleFooBG int

type (
	// godoc // want `godoc should start with symbol name \(pattern ""\)`
	TMultiFooBG int

	// godoc // want `godoc should start with symbol name \(pattern ""\)`
	TMultiBarBG int
)
