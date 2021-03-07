package cli

import (
	gitpkg "github.com/cbuschka/versioner/internal/git"
	"github.com/stretchr/testify/assert"
	"testing"
)

type GitMock struct {
	tags []string
}

func (gitMock GitMock) ListTags() ([]string, error) {
	return gitMock.tags, nil
}

func TestGetLatestVersionWhenNoTagsPresent(t *testing.T) {
	gitMock := GitMock{[]string{}}

	latestVersion, _ := getLatestVersion(gitpkg.Git(gitMock))

	assert.Nil(t, latestVersion)
}
