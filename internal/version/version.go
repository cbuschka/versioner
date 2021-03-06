package version

import (
	"fmt"
	goversion "github.com/hashicorp/go-version"
	"strings"
)

type Version struct {
	version *goversion.Version
}

func InitialVersion() *Version {
	version, err := ParseVersion("v1.0.0")
	if err != nil {
		panic(err)
	}

	return version
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

func (version Version) Next() *Version {
	prefix := ""
	if strings.HasPrefix(version.version.Original(), "v") {
		prefix = "v"
	}
	segments := version.version.Segments()
	if version.version.Prerelease() == "" {
		segments[len(segments)-1] = segments[len(segments)-1]+1
	}

	return toVersion(prefix, segments)
}

func toVersion(prefix string, segments []int) *Version {
	var buf strings.Builder
	fmt.Fprintf(&buf, "%s", prefix)
	for index, segment := range segments {
		if index == 0 {
			fmt.Fprintf(&buf, "%d", segment)
		} else {
			fmt.Fprintf(&buf, ".%d", segment)
		}
	}
	version, err := ParseVersion(buf.String())
	if err != nil {
		panic(err)
	}

	return version
}
