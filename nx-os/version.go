package version

import (
	"cmp"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// Version represents a NX-OS version
// https://sec.cloudapps.cisco.com/security/center/resources/ios_nx_os_reference_guide#release_naming_nx_os
type Version struct {
	Major               int
	Minor               int
	Maintenance         string
	Platform            string
	PlatformMinor       int
	PlatformMaintenance string
}

// NewVersion returns a parsed version
func NewVersion(ver string) (Version, error) {
	lhs, rhs, ok := strings.Cut(ver, ".")
	if !ok {
		return Version{}, fmt.Errorf("unexpected NX-OS version format. expected: %q, actual: %q", "<major>.<minor>\\(<maintenance>\\)(<platform>(<platform minor>)(<platform maintenance>))", ver)
	}
	major, err := strconv.Atoi(lhs)
	if err != nil {
		return Version{}, fmt.Errorf("parse major version. err: %w", err)
	}

	lhs, rhs, ok = strings.Cut(rhs, "(")
	if !ok {
		return Version{}, fmt.Errorf("unexpected NX-OS version format. expected: %q, actual: %q", "<major>.<minor>\\(<maintenance>\\)(<platform>(<platform minor>)(<platform maintenance>))", ver)
	}
	minor, err := strconv.Atoi(lhs)
	if err != nil {
		return Version{}, fmt.Errorf("parse minor version. err: %w", err)
	}

	lhs, rhs, ok = strings.Cut(rhs, ")")
	if !ok {
		return Version{}, fmt.Errorf("unexpected NX-OS version format. expected: %q, actual: %q", "<major>.<minor>\\(<maintenance>\\)(<platform>(<platform minor>)(<platform maintenance>))", ver)
	}
	maintenance := lhs

	if rhs == "" {
		return Version{Major: major, Minor: minor, Maintenance: maintenance}, nil
	}

	platform, platformMinor, platformMaintenance, err := func() (string, int, string, error) {
		lhs, rhs, ok = strings.Cut(rhs, "(")

		i := 0
		for ; i < len(lhs); i++ {
			if !unicode.IsUpper(rune(lhs[i])) {
				break
			}
		}

		var platformMinor int
		if lhs[i:] != "" {
			n, err := strconv.Atoi(lhs[i:])
			if err != nil {
				return "", 0, "", fmt.Errorf("parse platform minor version. err: %w", err)
			}
			platformMinor = n
		}

		if !ok {
			return lhs[:i], platformMinor, "", nil
		}
		return lhs[:i], platformMinor, strings.TrimSuffix(rhs, ")"), nil
	}()
	if err != nil {
		return Version{}, fmt.Errorf("parse platform part. err: %w", err)
	}

	return Version{
		Major:               major,
		Minor:               minor,
		Maintenance:         maintenance,
		Platform:            platform,
		PlatformMinor:       platformMinor,
		PlatformMaintenance: platformMaintenance,
	}, nil
}

var ErrCannotCompareDifferentPlatforms = fmt.Errorf("cannot compare versions with different platforms")

// Compare returns an integer comparing two version.
// The result will be 0 if v1==v2, -1 if v1 < v2, and +1 if v1 > v2.
func (v1 Version) Compare(v2 Version) (int, error) {
	if r := cmp.Or(
		cmp.Compare(v1.Major, v2.Major),
		cmp.Compare(v1.Minor, v2.Minor),
		cmp.Compare(v1.Maintenance, v2.Maintenance),
	); r != 0 {
		return r, nil
	}

	if v1.Platform != v2.Platform {
		return 0, ErrCannotCompareDifferentPlatforms
	}

	return cmp.Or(
		cmp.Compare(v1.PlatformMinor, v2.PlatformMinor),
		cmp.Compare(v1.PlatformMaintenance, v2.PlatformMaintenance),
	), nil
}

// String returns the full version string
func (v Version) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%d.%d(%s)", v.Major, v.Minor, v.Maintenance))
	if v.Platform != "" {
		sb.WriteString(v.Platform)
		if v.PlatformMinor > 0 {
			sb.WriteString(fmt.Sprintf("%d", v.PlatformMinor))
		}
		if v.PlatformMaintenance != "" {
			sb.WriteString(fmt.Sprintf("(%s)", v.PlatformMaintenance))
		}
	}
	return sb.String()
}
