// EXEMPT: this is a very very very very very very long line that would normally violate max-len //foo:bar // want `godoc line is too long \(67 > 40\)`
// another very very very long line for a variable that is [ALSO EXEMPT] from checks
// but only this line should be marked as long since it's not ignored.
package max_len

// EXEMPT: long package-level constant godoc that should be ignored by max-len //foo:bar // want `godoc line is too long \(67 > 40\)`
// another very very very long line for a variable that is [ALSO EXEMPT] from checks
// but only this line should be marked as long since it's not ignored.
const NoExemptConst = 0

// EXEMPT: long variable godoc that should be ignored by max-len //foo:bar // want `godoc line is too long \(67 > 40\)`
// another very very very long line for a variable that is [ALSO EXEMPT] from checks
// but only this line should be marked as long since it's not ignored.
var NoExemptVar = 0
