package version

import (
	"cmp"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// Version represents an IOS version
// https://sec.cloudapps.cisco.com/security/center/resources/ios_nx_os_reference_guide#release_naming_ios
type Version struct {
	Major       int
	Minor       int
	Feature     string
	Release     string
	Maintenance string
}

// NewVersion returns a parsed version
func NewVersion(ver string) (Version, error) {
	lhs, rhs, ok := strings.Cut(ver, ".")
	if !ok {
		return Version{}, fmt.Errorf("unexpected IOS version format. expected: %q, actual: %q", "<major>.<minor>\\(<feature>\\)(<release><maintenance>)", ver)
	}
	major, err := strconv.Atoi(lhs)
	if err != nil {
		return Version{}, fmt.Errorf("parse major version. err: %w", err)
	}

	lhs, rhs, ok = strings.Cut(rhs, "(")
	if !ok {
		return Version{}, fmt.Errorf("unexpected IOS version format. expected: %q, actual: %q", "<major>.<minor>\\(<feature>\\)(<release><maintenance>)", ver)
	}
	minor, err := strconv.Atoi(lhs)
	if err != nil {
		return Version{}, fmt.Errorf("parse minor version. err: %w", err)
	}

	lhs, rhs, ok = strings.Cut(rhs, ")")
	if !ok {
		return Version{}, fmt.Errorf("unexpected IOS version format. expected: %q, actual: %q", "<major>.<minor>\\(<feature>\\)(<release><maintenance>)", ver)
	}
	feature := lhs

	if rhs == "" {
		return Version{Major: major, Minor: minor, Feature: feature}, nil
	}

	release, maintenance := func() (string, string) {
		i := 0
		for ; i < len(rhs); i++ {
			if !unicode.IsUpper(rune(rhs[i])) {
				break
			}
		}
		return rhs[:i], rhs[i:]
	}()

	return Version{Major: major, Minor: minor, Feature: feature, Release: release, Maintenance: maintenance}, nil
}

var ErrCannotCompareDifferentRelease = fmt.Errorf("cannot compare versions with different release types")

// Compare returns an integer comparing two version.
// The result will be 0 if v1==v2, -1 if v1 < v2, and +1 if v1 > v2.
func (v1 Version) Compare(v2 Version) (int, error) {
	if r := cmp.Or(
		cmp.Compare(v1.Major, v2.Major),
		cmp.Compare(v1.Minor, v2.Minor),
		cmp.Compare(v1.Feature, v2.Feature),
	); r != 0 {
		return r, nil
	}

	if v1.Release != v2.Release {
		return 0, ErrCannotCompareDifferentRelease
	}
	return cmp.Compare(v1.Maintenance, v2.Maintenance), nil
}

// String returns the full version string
func (v Version) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%d.%d(%s)", v.Major, v.Minor, v.Feature))
	if v.Release != "" {
		sb.WriteString(v.Release)
	}
	if v.Maintenance != "" {
		sb.WriteString(v.Maintenance)
	}
	return sb.String()
}
