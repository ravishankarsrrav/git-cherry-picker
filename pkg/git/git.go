package git

import (
	"fmt"
	"github.com/fatih/color"
	"os/exec"
	"strings"
)

type Helper interface {
	CheckOut()
	CherryPick()
}

type GitHelper struct {
	FromBranch string
	ToBranch   string
}

/*CheckOut checks out the to branch*/
func (h *GitHelper) CheckOut() error {
	out, err := exec.Command("git", "checkout", h.ToBranch).CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
	}
	return err
}

/*ListCommits extracts the list of commits*/
func (h *GitHelper) ListCommits() ([]string, error) {
	out, err := exec.Command("git", "log", "--oneline", h.FromBranch).CombinedOutput()
	if err != nil {
		color.HiYellow(string(out))
		return []string{}, err
	}
	stringOut := string(out)
	splittedString := strings.Split(stringOut, "\n")
	splittedString = splittedString[:len(splittedString)-1]
	return splittedString, err
}

/*CherryPick cherry picks the commit ID*/
func (h *GitHelper) CherryPick(commitID string) error {
	out, err := exec.Command("git", "cherry-pick", commitID).CombinedOutput()
	if err != nil {
		color.HiYellow(string(out))
		return err
	}
	return err
}

/*AddAllChanges tracks and adds all the file changes*/
func (h *GitHelper) AddAllChanges() error {
	out, err := exec.Command("git", "add", ".").CombinedOutput()
	if err != nil {
		color.HiYellow(string(out))
	}
	return err
}

/*Continue continues the cherry-pick process*/
func (h *GitHelper) Continue() error {
	out, err := exec.Command("git", "cherry-pick", "--continue").CombinedOutput()
	if err != nil {
		color.HiYellow(string(out))
	}
	return err
}
