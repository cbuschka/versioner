package version

import (
	goversion "github.com/hashicorp/go-version"
)

//Collection is a wrapper to make Version sortable
type Collection []*goversion.Version

// Len gives length of collection.
func (v Collection) Len() int {
	return len(v)
}

// Less compares two version items.
func (v Collection) Less(i, j int) bool {
	return v[i].LessThan(v[j])
}

// Swap swaps two version items.
func (v Collection) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}
