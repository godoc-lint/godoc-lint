// Package version provides versioning information for the module.
package version

import "fmt"

// Current represents the current version.
var Current = Version{
	Major:  0,
	Minor:  10,
	Patch:  2,
	Suffix: "",
}

// Version represents module version (in semver format).
type Version struct {
	Major  uint
	Minor  uint
	Patch  uint
	Suffix string
}

// String returns the string representation of the current instance.
func (v Version) String() string {
	suffix := ""
	if v.Suffix != "" {
		suffix = "-" + v.Suffix
	}
	return fmt.Sprintf("%d.%d.%d%s", v.Major, v.Minor, v.Patch, suffix)
}
