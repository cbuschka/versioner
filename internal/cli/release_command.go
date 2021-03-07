package cli

import (
	gitpkg "github.com/cbuschka/versioner/internal/git"
)

// ReleaseCommandConfig is the config for the release command.
type ReleaseCommandConfig struct {
}

// Run executes the release command.
func (config *ReleaseCommandConfig) Run() error {

	git, err := gitpkg.GetGit()
	if err != nil {
		return err
	}

	err = release(git)

	return err
}
