// some header

// Foo is a fake command.
//
// This test is to ensure that the godoc for command packages (i.e. `main`
// packages) are not enforced to begin with "Package main".
//
// See for more details:
//   - https://github.com/godoc-lint/godoc-lint/issues/10
//   - https://go.dev/doc/comment#cmd
package main_test
