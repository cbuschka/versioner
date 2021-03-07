package cli

import (
	gitpkg "github.com/cbuschka/versioner/internal/git"
)

type LatestVersionCommandConfig struct {
}

func (config *LatestVersionCommandConfig) Run() error {

	git := gitpkg.NewGit()

	latestVersion, err := getLatestVersion(git)
	if err != nil {
		return err
	}

	if latestVersion != nil {
		Print("%s", latestVersion.Original())
	} else {
		Debug("No versions found.")
	}

	return nil
}
