package cli

import (
	gitpkg "github.com/cbuschka/versioner/internal/git"
	"github.com/cbuschka/versioner/internal/version"
)

func getNextVersion(git gitpkg.Git) (*version.Version, error) {

	latestVersion, err := getLatestVersion(git)
	if err != nil {
		return nil, err
	}

	var nextVersion *version.Version
	if latestVersion != nil {
		nextVersion = latestVersion.Next()
	} else {
		nextVersion = version.InitialVersion()
	}

	return nextVersion, nil
}
