package cli

import (
        gitpkg "github.com/cbuschka/versioner/internal/git"
        "github.com/cbuschka/versioner/internal/version"
)

type NextVersionCommandConfig struct {
}

func (config *NextVersionCommandConfig) Run() error {

	git := gitpkg.NewGit()

	latestVersion, err := getLatestVersion(git)
	if err != nil {
		return err
	}

	var nextVersion *version.Version
	if latestVersion != nil {
		nextVersion = latestVersion.Next()
	} else {
		nextVersion = version.InitialVersion()
	}
	Print("%s", nextVersion.Original())

	return nil
}
