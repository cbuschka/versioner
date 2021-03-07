package cli

import (
	gitpkg "github.com/cbuschka/versioner/internal/git"
)

// NextVersionCommandConfig is the config for the next-version command.
type NextVersionCommandConfig struct {
}

// Run executes the next-version command.
func (config *NextVersionCommandConfig) Run() error {

	git := gitpkg.NewGit()

	nextVersion, err := getNextVersion(git)
	if err != nil {
		return err
	}

	Print("%s", nextVersion.Original())

	return nil
}
