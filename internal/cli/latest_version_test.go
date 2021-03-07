package cli

import (
	gitpkg "github.com/cbuschka/versioner/internal/git"
	"github.com/stretchr/testify/assert"
	"testing"
)

type GitMock struct {
	tags           []string
	clean          bool
	recordedTags   []string
	recordedPushes []bool
}

func (gitMock *GitMock) ListTags() ([]string, error) {
	return gitMock.tags, nil
}

func (gitMock *GitMock) IsWorkspaceClean() (bool, error) {
	return gitMock.clean, nil
}

func (gitMock *GitMock) Push(withTags bool) error {
	gitMock.recordedPushes = append(gitMock.recordedPushes, withTags)
	return nil
}

func (gitMock *GitMock) Tag(tag string) error {
	gitMock.recordedTags = append(gitMock.recordedTags, tag)
	return nil
}

func TestGetLatestVersionWhenNoTagsPresent(t *testing.T) {
	gitMock := &GitMock{[]string{}, true, []string{}, []bool{}}

	latestVersion, _ := getLatestVersion(gitpkg.Git(gitMock))

	assert.Nil(t, latestVersion)
}
