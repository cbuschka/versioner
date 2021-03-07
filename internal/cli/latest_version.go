package cli

import (
	gitpkg "github.com/cbuschka/versioner/internal/git"
	"github.com/cbuschka/versioner/internal/version"
)

func getLatestVersion(git gitpkg.Git) (*version.Version, error) {

	tags, err := git.ListTags()
	if err != nil {
		return nil, err
	}
	Debug("Tags found: %v\n", tags)

	versions, removedTags, err := version.FilterConvertAndSortAsc(tags)
	if err != nil {
		return nil, err
	}
	Debug("Invalid tags skipped: %v\n", removedTags)
	Debug("Sorted versions: %v\n", versions)

	if len(versions) > 0 {
		return versions[len(versions)-1], nil
	}

	return nil, nil
}
