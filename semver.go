package semver

import (
	"fmt"
	"strconv"
	"strings"
)

type Version struct {
	Major int
	Minor int
	Patch int
}

func NewVersion(major, minor, patch int) *Version {
	return &Version{
		Major: major,
		Minor: minor,
		Patch: patch,
	}
}

func (r *Version) UnmarshalJSON(bs []byte) error {
	ver := loadVersion(string(bs))
	if ver.IsEmpty() {
		return nil
	}

	*r = *ver
	return nil
}

func (r Version) MarshalJSON() ([]byte, error) {
	return []byte(`"` + r.String() + `"`), nil
}

func (r *Version) BiggerEqualThan(b *Version) bool {
	if r == nil {
		return b == nil
	}
	if b == nil {
		return true
	}

	if r.Major != b.Major {
		return r.Major > b.Major
	}
	if r.Minor != b.Minor {
		return r.Minor > b.Minor
	}
	return r.Patch >= b.Patch
}

func (r *Version) AddPatch(v int) *Version {
	return &Version{
		Major: r.Major,
		Minor: r.Minor,
		Patch: r.Patch + v,
	}
}

func (r *Version) String() string {
	return fmt.Sprintf("v%d.%d.%d", r.Major, r.Minor, r.Patch)
}

func (r *Version) IsEmpty() bool {
	return r.Major == 0 && r.Minor == 0 && r.Patch == 0
}

func loadVersion(std string) *Version {
	std = strings.TrimLeft(std, `"`)
	std = strings.TrimRight(std, `"`)
	std = strings.TrimPrefix(strings.ToLower(std), "v")

	stds := strings.Split(std, ".") // std-version

	var (
		major int64
		minor int64
		patch int64
	)
	if len(stds) >= 1 {
		major, _ = strconv.ParseInt(stds[0], 10, 64)
	}
	if len(stds) >= 2 {
		minor, _ = strconv.ParseInt(stds[1], 10, 64)
	}
	if len(stds) >= 3 {
		patch, _ = strconv.ParseInt(stds[2], 10, 64)
	}

	return &Version{
		Major: int(major),
		Minor: int(minor),
		Patch: int(patch),
	}
}
