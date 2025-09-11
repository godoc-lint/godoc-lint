// some header

package deprecated

// some godoc
//
// Deprecated: do not use
const Foo = 0

// Deprecated: do not use
//
// some godoc
const Bar = 0

// some godoc
//
// Deprecated: do not use
const (
	Baz = 0
)

// some godoc
const (
	// some godoc
	//
	// Deprecated: do not use
	Yo = 0
)

// some godoc
//
// Deprecated: do not use
const (
	// some godoc
	//
	// Deprecated: do not use
	Yolo = 0
)

// The //foo:bar directives mark the trailing comment as a directive so they're
// not parsed as a normal trailing comment group.

// Deprecated: //foo:bar
type Alpha int

// Deprecated: do not use
type Beta int

// Deprecated:  do not use
type Charlie int

// Deprecated:  deprecated: do not use
type Delta int

// this is a symbol but here are the reasons why it's
// deprecated: blah, blah, ...
func Echo() {}

// this is a symbol
// deprecated: do not use
func Foxtrot() {}

// DEPRECATED: bad deprecation note but okay since the symbol is unexported
func golf() {}
