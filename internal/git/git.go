package git

import (
	"os/exec"
	"strings"
)

type Git interface {
	ListTags() ([]string, error)
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
