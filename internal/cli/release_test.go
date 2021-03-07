package cli

import (
	gitpkg "github.com/cbuschka/versioner/internal/git"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReleaseWhenNoTagsPresent(t *testing.T) {
	gitMock := &GitMock{[]string{}, true, []string{}, []bool{}}

	err := release(gitpkg.Git(gitMock))

	assert.Nil(t, err)
	assert.Equal(t, []string{"v1.0.0"}, gitMock.recordedTags)
	assert.Equal(t, []bool{true}, gitMock.recordedPushes)
}
