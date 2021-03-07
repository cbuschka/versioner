package git

import (
	"os/exec"
	"strings"
)

type Git interface {
	ListTags() ([]string, error)
	IsWorkspaceClean() (bool, error)
}

type GitSession struct {}

func NewGit() *GitSession {
	return &GitSession{}
}

// ListTags calls git command and lists tags
func (git *GitSession) ListTags() ([]string, error) {
	out, err := exec.Command("git", "tag", "--list").Output()
	if err != nil {
		return nil, err
	}
	outText := string(out)
	tags := strings.Split(outText, "\n")
	return tags, nil
}

func (git *GitSession) IsWorkspaceClean() (bool, error) {
	out, err := exec.Command("git", "status", "--porcelain").Output()
	if err != nil {
		return false, err
	}
	isClean := string(out) == ""
	return isClean, nil
}
