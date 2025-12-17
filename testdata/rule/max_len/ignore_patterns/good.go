// EXEMPT: this is a very very very very very very long line that would normally violate max-len
// another very very very long line for a variable that is [ALSO EXEMPT] from checks
package max_len

// EXEMPT: long package-level constant godoc that should be ignored by max-len
// another very very very long line for a variable that is [ALSO EXEMPT] from checks
// And this line is okay in length.
const ExemptConst = 0

// EXEMPT: long variable godoc that should be ignored by max-len
// another very very very long line for a variable that is [ALSO EXEMPT] from checks
// And this line is okay in length.
var ExemptVar = 0
