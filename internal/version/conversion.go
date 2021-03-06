package version

import (
    "log"
    goversion "github.com/hashicorp/go-version"
    "strings"
    "sort"
)

func ConvertAndSortAsc(versionStrs []string) ([]*goversion.Version, error) {
    versions, err := convert(versionStrs)
    if err != nil {
      return nil, err
    }

    sort.Sort(VersionCollection(versions))
    return versions, nil
}

func convert(versionStrs []string) ([]*goversion.Version, error) {
    versions := make([]*goversion.Version,0)
    for _, versionStr := range versionStrs {
        versionStr = strings.TrimSpace(versionStr)
        if len(versionStr) > 0 {
          version, err := goversion.NewVersion(versionStr)
          if err != nil {
              log.Printf("err %v with %s\n", err, versionStr)
          } else {
              versions = append(versions, version)
          }
        }
    }
    return versions, nil
}
