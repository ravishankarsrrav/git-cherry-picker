package git

import (
	"fmt"
	"os/exec"
	"strings"
)

type Helper interface {
	CheckOut()
}

type GitHelper struct {
	FromBranch string
	ToBranch   string
}

func (h *GitHelper) CheckOut() error {
	out, err := exec.Command("git", "checkout", h.ToBranch).Output()
	fmt.Println(out)
	return err
}

func (h *GitHelper) ListCommits() ([]string, error) {
	out, err := exec.Command("git", "log", "--oneline").Output()
	if err != nil {
		return []string{}, err
	}
	stringOut := string(out)
	splittedString := strings.Split(stringOut, "\n")
	splittedString = splittedString[:len(splittedString)-1]
	return splittedString, err
}
