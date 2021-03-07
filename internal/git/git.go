package git

import (
	"os"
	"os/exec"
	"strings"
)

// Git is the interface to a git session.
type Git interface {
	ListTags() ([]string, error)
	IsWorkspaceClean() (bool, error)
	Tag(tags string) error
	Push(withTags bool) error
}

// Session is a git session.
type Session struct {
	repoDir string
}

// GetGit produces a new git session.
func GetGit() (*Session, error) {
	workingDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	return &Session{workingDir}, nil
}

// NewGit creates a new git session with repo at repoDir.
func NewGit(repoDir string) (*Session, error) {
	err := os.Chdir(repoDir)
	if err != nil {
		return nil, err
	}

	return &Session{repoDir}, err
}

// ListTags calls git command and lists tags
func (git *Session) ListTags() ([]string, error) {

	err := os.Chdir(git.repoDir)
	if err != nil {
		return nil, err
	}

	out, err := exec.Command("git", "tag", "--list").Output()
	if err != nil {
		return nil, err
	}
	outText := string(out)
	tags := strings.Split(outText, "\n")
	return tags, nil
}

// IsWorkspaceClean checks if the workspace is clean.
func (git *Session) IsWorkspaceClean() (bool, error) {
	err := os.Chdir(git.repoDir)
	if err != nil {
		return false, err
	}

	out, err := exec.Command("git", "status", "--porcelain").Output()
	if err != nil {
		return false, err
	}
	isClean := string(out) == ""
	return isClean, nil
}

// Tag creates a tag.
func (git *Session) Tag(tag string) error {
	err := os.Chdir(git.repoDir)
	if err != nil {
		return err
	}

	_, err = exec.Command("git", "tag", tag).Output()
	if err != nil {
		return err
	}
	return nil
}

// Push pushes the git repository status to origin
func (git *Session) Push(withTags bool) error {
	err := os.Chdir(git.repoDir)
	if err != nil {
		return err
	}

	_, err = exec.Command("git", "push", "--tags", "--porcelain").Output()
	if err != nil {
		return err
	}
	return nil
}
