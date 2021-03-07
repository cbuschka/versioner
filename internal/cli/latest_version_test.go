package cli

import (
	gitpkg "github.com/cbuschka/versioner/internal/git"
	"github.com/stretchr/testify/assert"
	"testing"
)

type GitMock struct {
	tags []string
	clean bool
}

func (gitMock GitMock) ListTags() ([]string, error) {
	return gitMock.tags, nil
}

func (gitMock GitMock) IsWorkspaceClean() (bool, error) {
	return gitMock.clean, nil
}

func TestGetLatestVersionWhenNoTagsPresent(t *testing.T) {
	gitMock := GitMock{[]string{},true}

	latestVersion, _ := getLatestVersion(gitpkg.Git(gitMock))

	assert.Nil(t, latestVersion)
}
