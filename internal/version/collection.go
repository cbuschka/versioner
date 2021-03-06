package version

import (
  goversion "github.com/hashicorp/go-version"
)

type VersionCollection []*goversion.Version

// Len gives length of collection.
func (v VersionCollection) Len() int {
	return len(v)
}

// Less compares two version items.
func (v VersionCollection) Less(i, j int) bool {
	return v[i].LessThan(v[j])
}

// Swap swaps two version items.
func (v VersionCollection) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}
