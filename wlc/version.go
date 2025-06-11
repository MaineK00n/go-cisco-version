package version

import (
	"cmp"
	"fmt"
	"strconv"
	"strings"
)

// Version represents an Cisco Wireless LAN Controller (WLC) version
type Version struct {
	Major       int
	Minor       int
	Maintenance int
	Build       int
}

// NewVersion returns a parsed version
func NewVersion(ver string) (Version, error) {
	switch ss := strings.Split(strings.TrimSuffix(strings.NewReplacer("(", ".", ")", ".").Replace(ver), "."), "."); len(ss) {
	case 3:
		major, err := strconv.Atoi(ss[0])
		if err != nil {
			return Version{}, fmt.Errorf("parse major version. err: %w", err)
		}

		minor, err := strconv.Atoi(ss[1])
		if err != nil {
			return Version{}, fmt.Errorf("parse minor version. err: %w", err)
		}

		maintenance, err := strconv.Atoi(ss[2])
		if err != nil {
			return Version{}, fmt.Errorf("parse maintenance version. err: %w", err)
		}

		return Version{Major: major, Minor: minor, Maintenance: maintenance}, nil
	case 4:
		major, err := strconv.Atoi(ss[0])
		if err != nil {
			return Version{}, fmt.Errorf("parse major version. err: %w", err)
		}

		minor, err := strconv.Atoi(ss[1])
		if err != nil {
			return Version{}, fmt.Errorf("parse minor version. err: %w", err)
		}

		maintenance, err := strconv.Atoi(ss[2])
		if err != nil {
			return Version{}, fmt.Errorf("parse maintenance version. err: %w", err)
		}

		build, err := strconv.Atoi(ss[3])
		if err != nil {
			return Version{}, fmt.Errorf("parse build version. err: %w", err)
		}

		return Version{Major: major, Minor: minor, Maintenance: maintenance, Build: build}, nil
	default:
		return Version{}, fmt.Errorf("unexpected WLC version format. expected: %q, actual: %q", "<major>.<minor>.<maintenance>(.<build>)", ver)
	}
}

// Compare returns an integer comparing two version.
// The result will be 0 if v1==v2, -1 if v1 < v2, and +1 if v1 > v2.
func (v1 Version) Compare(v2 Version) int {
	return cmp.Or(
		cmp.Compare(v1.Major, v2.Major),
		cmp.Compare(v1.Minor, v2.Minor),
		cmp.Compare(v1.Maintenance, v2.Maintenance),
		cmp.Compare(v1.Build, v2.Build),
	)
}

// String returns the full version string
func (v Version) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", v.Major, v.Minor, v.Maintenance, v.Build)
}
