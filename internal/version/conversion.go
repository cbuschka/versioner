package version

import (
	"sort"
	"strings"
)

// FilterConvertAndSortAsc sorts versionStrs matching the semver format in ascending order and
// returns removedVersions as strings in input order
func FilterConvertAndSortAsc(versionStrs []string) ([]*Version, []string, error) {
	versions, removedVersionStrs, err := convert(versionStrs)
	if err != nil {
		return nil, nil, err
	}

	sort.Sort(Collection(versions))
	return versions, removedVersionStrs, nil
}

func convert(versionStrs []string) ([]*Version, []string, error) {
	versions := make([]*Version, 0)
	removedVersionStrs := make([]string, 0)
	for _, versionStr := range versionStrs {
		versionStr = strings.TrimSpace(versionStr)
		if len(versionStr) > 0 {
			version, err := ParseVersion(versionStr)
			if err != nil {
				removedVersionStrs = append(removedVersionStrs, versionStr)
			} else {
				versions = append(versions, version)
			}
		}
	}
	return versions, removedVersionStrs, nil
}
