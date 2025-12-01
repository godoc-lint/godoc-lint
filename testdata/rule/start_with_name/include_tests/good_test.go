package include_tests

// TestFoo is anything.
const TestFoo = 0

// TestOwl is a bird.
var TestOwl = 0

// TestCat is asleep.
type TestCat int

// TESTYolo has a godoc.
func TESTYolo() {}

// testFoo is anything.
const testFoo = 0

// testOwl is a bird.
var testOwl = 0

// testCat is asleep.
type testCat int

// testYolo has a godoc.
func testYolo() {}

// Bad godoc, but should be ignored due to blank identifier.
var _ = 0
