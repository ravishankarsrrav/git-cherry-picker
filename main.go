package main

import (
	"fmt"
	"github.com/ravishankarsrrav/git-cherry-picker/pkg/git"
	"github.com/ravishankarsrrav/git-cherry-picker/pkg/ui"
	flag "github.com/spf13/pflag"
)

var fromBranch string
var toBranch string

func main() {
	flag.StringVar(&fromBranch, "from_branch", "", "name of the branch from which the commit is cherry picked")
	flag.StringVar(&toBranch, "to_branch", "main", "name of the branch to which the commit is cherry picked")
	flag.Parse()
	if toBranch == "" {
		fmt.Println("toBranch should be provided")
	}
	var gitHelper = git.GitHelper{FromBranch: fromBranch, ToBranch: toBranch}
	/*err := gitHelper.CheckOut()
	if err != nil {
		fmt.Printf("checkout failed %s\n", err.Error())
	}*/
	commits, err := gitHelper.ListCommits()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(commits)
	selectedElement := ui.Draw(commits)
	print(selectedElement)

}
