package include_tests_test

// (BG: bad godoc)

// godoc // want `godoc should start with symbol name \(pattern "%"\)`
const TestFooBG = 0

// godoc // want `godoc should start with symbol name \(pattern "%"\)`
var TestOwlBG = 0

// godoc // want `godoc should start with symbol name \(pattern "%"\)`
type TestCatBG int

// godoc // want `godoc should start with symbol name \(pattern "%"\)`
func TESTYoloBG() {}

// godoc // want `godoc should start with symbol name \(pattern "%"\)`
const testFooBG = 0

// godoc // want `godoc should start with symbol name \(pattern "%"\)`
var testOwlBG = 0

// godoc // want `godoc should start with symbol name \(pattern "%"\)`
type testCatBG int

// godoc // want `godoc should start with symbol name \(pattern "%"\)`
func testYoloBG() {}
