package cli

import (
	gitpkg "github.com/cbuschka/versioner/internal/git"
)

type ReleaseCommandConfig struct {
}

func (config *ReleaseCommandConfig) Run() error {

	git := gitpkg.NewGit()

	err := release(git)

	return err
}
