package git

import (
	"os/exec"
	"strings"
)

// ListTags calls git command and lists tags
func ListTags() ([]string, error) {
	out, err := exec.Command("git", "tag", "--list").Output()
	if err != nil {
		return nil, err
	}
	outText := string(out)
	tags := strings.Split(outText, "\n")
	return tags, nil
}
