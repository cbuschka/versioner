package cli

import (
	gitpkg "github.com/cbuschka/versioner/internal/git"
)

// LatestVersionCommandConfig is the command config for latest-version command.
type LatestVersionCommandConfig struct {
}

// Run executes the latest-version command.
func (config *LatestVersionCommandConfig) Run() error {

	git, err := gitpkg.GetGit()
	if err != nil {
		return err
	}

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
