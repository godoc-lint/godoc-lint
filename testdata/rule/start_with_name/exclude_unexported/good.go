package strict

// (BG: bad godoc)

// godoc
const fooBG = 0

// godoc
var owlBG = 0

// godoc
type catBG int

// godoc
func yoloBG() {}

// Bad godoc, but should be ignored due to blank identifier.
var _ = 0

// Bad godoc but should be ignored since the receiver base type is unexported.
func (catBG) Foo() {}
