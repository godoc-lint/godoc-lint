// some header

// godoc with multiple unused links // want `godoc has unused link \("link1"\)` `godoc has unused link \("link2"\)`
//
// [link1]: https://foo.com
// [link2]: https://foo.com
package no_unused_link

// godoc with multiple unused links // want `godoc has unused link \("link1"\)` `godoc has unused link \("link2"\)`
//
// [link1]: https://foo.com
// [link2]: https://foo.com
const MultipleUnusedLinkConst = 0

// godoc with multiple unused links // want `godoc has unused link \("link1"\)` `godoc has unused link \("link2"\)`
//
// [link1]: https://foo.com
// [link2]: https://foo.com
const (
	// godoc with multiple unused links // want `godoc has unused link \("link1"\)` `godoc has unused link \("link2"\)`
	//
	// [link1]: https://foo.com
	// [link2]: https://foo.com
	MultiMultipleUnusedLinkConst = 0
)

// godoc with multiple unused links // want `godoc has unused link \("link1"\)` `godoc has unused link \("link2"\)`
//
// [link1]: https://foo.com
// [link2]: https://foo.com
type TMultipleUnusedLink int

// godoc with multiple unused links // want `godoc has unused link \("link1"\)` `godoc has unused link \("link2"\)`
//
// [link1]: https://foo.com
// [link2]: https://foo.com
type (
	TMultiMultipleUnusedLink int
)

// godoc with multiple unused links // want `godoc has unused link \("link1"\)` `godoc has unused link \("link2"\)`
//
// [link1]: https://foo.com
// [link2]: https://foo.com
func MultipleUnusedLinkFunc() {}
