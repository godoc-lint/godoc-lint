// EXEMPT: this is a very very very very very very long line that would normally violate max-len
// another very very very long line for a variable that is [ALSO EXEMPT] from checks
// but only this line should be marked as long since it's not ignored.  // want `godoc line is too long \(114 > 40\)`
package max_len

// EXEMPT: long package-level constant godoc that should be ignored by max-len
// another very very very long line for a variable that is [ALSO EXEMPT] from checks
// but only this line should be marked as long since it's not ignored.  // want `godoc line is too long \(114 > 40\)`
const NoExemptConst = 0

// EXEMPT: long variable godoc that should be ignored by max-len
// another very very very long line for a variable that is [ALSO EXEMPT] from checks
// but only this line should be marked as long since it's not ignored.  // want `godoc line is too long \(114 > 40\)`
var NoExemptVar = 0
