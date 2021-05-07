package git

import "os/exec"

type Helper interface {
	CheckOut()
}

type GitHelper struct {
	FromBranch string
	ToBranch   string
}

func (h *GitHelper) CheckOut() error {
	_, err := exec.Command("git", "checkout", h.ToBranch).Output()
	return err
}
