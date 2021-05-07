package git

import "os/exec"

type Helper interface {
	CheckOut()
}

type helper struct {
	fromBranch string
	toBranch   string
}

func (h helper) CheckOut() error {
	_, err := exec.Command("git", "checkout", h.toBranch).Output()
	return err
}
