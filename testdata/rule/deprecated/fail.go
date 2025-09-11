// some header

// (BG: bad godoc)

// some godoc // want `deprecation note should be formatted as "Deprecated: "`
//
// deprecated: do not use
//
// Deprecated: valid deprecation
package deprecated

// deprecated: do not use // want `deprecation note should be formatted as "Deprecated: "`
//
// godoc
const FooBG = 0

// some godoc // want `deprecation note should be formatted as "Deprecated: "`
//
// deprecated:do not use
const (
	BarBG = 0
)

// some godoc // want `deprecation note should be formatted as "Deprecated: "`
//
// deprecated:: do not use
const (
	// some godoc
	BazBG = 0
)

// some godoc // want `deprecation note should be formatted as "Deprecated: "`
//
// DEPRECATED: do not use
const (
	// some godoc // want `deprecation note should be formatted as "Deprecated: "`
	//
	// deprecated:?do not use
	YoBG = 0
)

// some godoc // want `deprecation note should be formatted as "Deprecated: "`
//
// DEPRECATED: do not use
const (
	// some godoc // want `deprecation note should be formatted as "Deprecated: "`
	//
	// DEPRECATED: do not use
	YoloBG = 0
)

// some godoc // want `deprecation note should be formatted as "Deprecated: "`
//
// DEPRECATED:do not use
type AlphaBG int

// some godoc // want `deprecation note should be formatted as "Deprecated: "`
//
// Deprecated:
// do not use
type BravoBG int

// some godoc // want `deprecation note should be formatted as "Deprecated: "`
//
// Deprecated:
type CharlieBG int

// some godoc // want `deprecation note should be formatted as "Deprecated: "`
//
// DEPRECATED:
// do not use
type DeltaBG int

// some godoc // want `deprecation note should be formatted as "Deprecated: "`
//
// DEPRECATED:
type EchoBG int

// The //foo:bar directives mark the trailing comment as a directive so they're
// not parsed as a normal trailing comment group.

// DEPRECATED://foo:bar // want `deprecation note should be formatted as "Deprecated: "`
type FoxtrotBG int

// DEPRECATED: //foo:bar // want `deprecation note should be formatted as "Deprecated: "`
type GolfBG int

// Deprecated://foo:bar // want `deprecation note should be formatted as "Deprecated: "`
type HotelBG int

// some godoc // want `deprecation note should be formatted as "Deprecated: "`
//
// DePREcatED: do not use
type IndiaBG int
