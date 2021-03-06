package version

import (
	goversion "github.com/hashicorp/go-version"
)

type Version struct {
	version *goversion.Version
}

func ParseVersion(raw string) (*Version, error) {
	version, err := goversion.NewVersion(raw)
	if err != nil {
		return nil, err
	}

	return &Version{version}, nil
}

func (version Version) String() string {
	return version.version.String()
}

func (version Version) Original() string {
	return version.version.Original()
}
