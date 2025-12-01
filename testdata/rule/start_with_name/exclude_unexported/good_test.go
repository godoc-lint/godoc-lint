package strict

// (BG: bad godoc)

// godoc
const testFooBG = 0

// godoc
var testOwlBG = 0

// godoc
type testCatBG int

// godoc
func testYoloBG() {}

// Bad godoc, but should be ignored due to blank identifier.
var _ = 0
