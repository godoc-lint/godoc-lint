// some header

// godoc with [link1] and [link2]
//
// [link1]: https://foo.com
// [link2]: https://foo.com
package no_unused_link

// godoc with [link1] and [link2]
//
// [link1]: https://foo.com
// [link2]: https://foo.com
const UsedLinkConst = 0

// godoc with [link1] and [link2]
//
// [link1]: https://foo.com
// [link2]: https://foo.com
const (
	// godoc with [link1] and [link2]
	//
	// [link1]: https://foo.com
	// [link2]: https://foo.com
	MultiUsedLinkConst = 0
)

// godoc with [link1] and [link2]
//
// [link1]: https://foo.com
// [link2]: https://foo.com
type TUsedLink int

// godoc with [link1] and [link2]
//
// [link1]: https://foo.com
// [link2]: https://foo.com
type (
	TMultiUsedLink int
)

// godoc with [link1] and [link2]
//
// [link1]: https://foo.com
// [link2]: https://foo.com
func UsedLinkFunc() {}
