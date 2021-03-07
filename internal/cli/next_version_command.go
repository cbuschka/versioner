package cli

import (
	gitpkg "github.com/cbuschka/versioner/internal/git"
)

type NextVersionCommandConfig struct {
}

func (config *NextVersionCommandConfig) Run() error {

	git := gitpkg.NewGit()

	nextVersion, err := getNextVersion(git)
	if err != nil {
		return err
	}

	Print("%s", nextVersion.Original())

	return nil
}
