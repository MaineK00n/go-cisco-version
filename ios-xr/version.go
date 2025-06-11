package version

import (
	"cmp"
	"fmt"
	"strconv"
	"strings"
)

// Version represents an IOS-XR version
// https://sec.cloudapps.cisco.com/security/center/resources/ios_nx_os_reference_guide#release_naming_ios_xr
// https://www.cisco.com/c/en/us/products/collateral/ios-nx-os-software/ios-xr-software/ios-xr-release-taxonomy-bulletin.html
type Version struct {
	Major   int
	Minor   int
	Release int
}

// NewVersion returns a parsed version
func NewVersion(ver string) (Version, error) {
	switch ss := strings.Split(ver, "."); len(ss) {
	case 3:
		major, err := strconv.Atoi(ss[0])
		if err != nil {
			return Version{}, fmt.Errorf("parse major version. err: %w", err)
		}

		minor, err := strconv.Atoi(ss[1])
		if err != nil {
			return Version{}, fmt.Errorf("parse minor version. err: %w", err)
		}

		release, err := strconv.Atoi(ss[2])
		if err != nil {
			return Version{}, fmt.Errorf("parse release version. err: %w", err)
		}

		return Version{
			Major:   major,
			Minor:   minor,
			Release: release,
		}, nil
	default:
		return Version{}, fmt.Errorf("unexpected IOS XR version format. expected: %q, actual: %q", "<major|year>.<minor|quarter>.<release>", ver)
	}
}

// Compare returns an integer comparing two version.
// The result will be 0 if v1==v2, -1 if v1 < v2, and +1 if v1 > v2.
func (v1 Version) Compare(v2 Version) int {
	return cmp.Or(
		cmp.Compare(v1.Major, v2.Major),
		cmp.Compare(v1.Minor, v2.Minor),
		cmp.Compare(v1.Release, v2.Release),
	)
}

// String returns the full version string
func (v Version) String() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Release)
}
