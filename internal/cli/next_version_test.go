package cli

import (
	gitpkg "github.com/cbuschka/versioner/internal/git"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetNextVersionWhenNoTagsPresent(t *testing.T) {
	gitMock := &GitMock{[]string{}, true, []string{}, []bool{}}

	nextVersion, _ := getNextVersion(gitpkg.Git(gitMock))

	assert.Equal(t, "v1.0.0", nextVersion.Original())
}

func TestGetNextVersionWhenTagsPresent(t *testing.T) {
	gitMock := &GitMock{[]string{"v999.2.0"}, true, []string{}, []bool{}}

	nextVersion, _ := getNextVersion(gitpkg.Git(gitMock))

	assert.Equal(t, "v999.2.1", nextVersion.Original())
}
