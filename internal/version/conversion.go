package version

import (
	goversion "github.com/hashicorp/go-version"
	"sort"
	"strings"
)

func FilterConvertAndSortAsc(versionStrs []string) ([]*goversion.Version, []string, error) {
	versions, removedVersionStrs, err := convert(versionStrs)
	if err != nil {
		return nil, nil, err
	}

	sort.Sort(Collection(versions))
	return versions, removedVersionStrs, nil
}

func convert(versionStrs []string) ([]*goversion.Version, []string, error) {
	versions := make([]*goversion.Version, 0)
	removedVersionStrs := make([]string, 0)
	for _, versionStr := range versionStrs {
		versionStr = strings.TrimSpace(versionStr)
		if len(versionStr) > 0 {
			version, err := goversion.NewVersion(versionStr)
			if err != nil {
				removedVersionStrs = append(removedVersionStrs, versionStr)
			} else {
				versions = append(versions, version)
			}
		}
	}
	return versions, removedVersionStrs, nil
}
