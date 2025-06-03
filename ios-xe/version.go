package version

import (
	"cmp"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// Version represents an IOS-XE version
// https://sec.cloudapps.cisco.com/security/center/resources/ios_nx_os_reference_guide#release_naming_ios_xe
type Version struct {
	Release     string
	Major       int
	Minor       int
	Maintenance string
}

// NewVersion returns a parsed version
func NewVersion(ver string) (Version, error) {
	switch ss := strings.Split(ver, "."); len(ss) {
	case 3:
		switch {
		case strings.HasPrefix(ss[0], "03"), strings.HasPrefix(ss[0], "3"):
			major, err := strconv.Atoi(ss[0])
			if err != nil {
				return Version{}, fmt.Errorf("parse major version. err: %w", err)
			}

			minor, err := strconv.Atoi(ss[1])
			if err != nil {
				return Version{}, fmt.Errorf("parse minor version. err: %w", err)
			}

			maintenance, release := func() (string, string) {
				i := 0
				for ; i < len(ss[2]); i++ {
					if unicode.IsUpper(rune(ss[2][i])) {
						break
					}
				}
				return ss[2][:i], ss[2][i:]
			}()

			return Version{
				Release:     release,
				Major:       major,
				Minor:       minor,
				Maintenance: maintenance,
			}, nil
		default:
			release, major, err := func() (string, int, error) {
				lhs, rhs, ok := strings.Cut(ss[0], "-")
				if ok {
					major, err := strconv.Atoi(rhs)
					if err != nil {
						return "", 0, fmt.Errorf("parse major version. err: %w", err)
					}
					return lhs, major, nil
				}

				major, err := strconv.Atoi(ss[0])
				if err != nil {
					return "", 0, fmt.Errorf("parse major version. err: %w", err)
				}
				return "", major, nil
			}()
			if err != nil {
				return Version{}, fmt.Errorf("parse release+major version. err: %w", err)
			}

			minor, err := strconv.Atoi(ss[1])
			if err != nil {
				return Version{}, fmt.Errorf("parse minor version. err: %w", err)
			}

			return Version{
				Release:     release,
				Major:       major,
				Minor:       minor,
				Maintenance: ss[2],
			}, nil
		}
	default:
		return Version{}, fmt.Errorf("unexpected IOS XE version format. expected: %q, actual: %q", []string{"(<release>-)<major>.<minor>.<maintenance>", "<major>.<minor>.<maintenance>(<release>)"}, ver)
	}
}

var ErrCannotCompareDifferentRelease = fmt.Errorf("cannot compare versions with different release types")

// Compare returns an integer comparing two version.
// The result will be 0 if v1==v2, -1 if v1 < v2, and +1 if v1 > v2.
func (v1 Version) Compare(v2 Version) (int, error) {
	if r := cmp.Compare(v1.Major, v2.Major); r != 0 {
		return r, nil
	}

	if v1.Major == 3 && v1.Release != v2.Release {
		return 0, ErrCannotCompareDifferentRelease
	}

	return cmp.Or(
		cmp.Compare(v1.Minor, v2.Minor),
		cmp.Compare(v1.Maintenance, v2.Maintenance),
	), nil
}

// String returns the full version string
func (v Version) String() string {
	switch v.Major {
	case 3:
		return fmt.Sprintf("%d.%d.%s%s", v.Major, v.Minor, v.Maintenance, v.Release)
	default:
		var sb strings.Builder
		if v.Release != "" {
			sb.WriteString(fmt.Sprintf("%s-", v.Release))
		}
		sb.WriteString(fmt.Sprintf("%d.%d.%s", v.Major, v.Minor, v.Maintenance))
		return sb.String()
	}
}
