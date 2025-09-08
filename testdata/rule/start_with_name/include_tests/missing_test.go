package include_tests

// (BG: bad godoc)

// godoc // want `godoc should start with symbol name \("TestFooBG"\)`
const TestFooBG = 0

// godoc // want `godoc should start with symbol name \("TestOwlBG"\)`
var TestOwlBG = 0

// godoc // want `godoc should start with symbol name \("TestCatBG"\)`
type TestCatBG int

// godoc // want `godoc should start with symbol name \("TESTYoloBG"\)`
func TESTYoloBG() {}

// godoc // want `godoc should start with symbol name \("testFooBG"\)`
const testFooBG = 0

// godoc // want `godoc should start with symbol name \("testOwlBG"\)`
var testOwlBG = 0

// godoc // want `godoc should start with symbol name \("testCatBG"\)`
type testCatBG int

// godoc // want `godoc should start with symbol name \("testYoloBG"\)`
func testYoloBG() {}
