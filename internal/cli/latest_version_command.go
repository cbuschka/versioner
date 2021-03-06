package cli

import (
	"github.com/cbuschka/versioner/internal/git"
	"github.com/cbuschka/versioner/internal/version"
)

type LatestVersionConfig struct {
}

func (config *LatestVersionConfig) Run() error {

	tags, err := git.ListTags()
	if err != nil {
		return err
	}
	Debug("Tags found: %v\n", tags)

	versions, removedTags, err := version.FilterConvertAndSortAsc(tags)
	if err != nil {
		return err
	}
	Debug("Invalid tags skipped: %v\n", removedTags)
	Debug("Sorted versions: %v\n", versions)

	if len(versions) > 0 {
		Print("%s", versions[len(versions)-1])
	}

	return nil
}
