package cli

import (
	gitpkg "github.com/cbuschka/versioner/internal/git"
)

func release(git gitpkg.Git) error {

	nextVersion, err := getNextVersion(git)
	if err != nil {
		return err
	}

	Print("Tagging as %s...", nextVersion.Original())
	err = git.Tag(nextVersion.Original())
	if err != nil {
		return err
	}

	Print("Pushing tags...")
	err = git.Push(true)
	if err != nil {
		return err
	}

	Print("Release successful.")

	return nil
}
