package main

import (
	"./pkg/git"
	"fmt"
	flag "github.com/spf13/pflag"
)

var fromBranch string
var toBranch string

func main() {
	flag.StringVar(&fromBranch, "from_branch", "", "name of the branch from which the commit is cherry picked")
	flag.StringVar(&toBranch, "to_branch", "", "name of the branch to which the commit is cherry picked")
	if fromBranch == "" {
		fmt.Println("fromBranch should be provided")
	}
	if toBranch == "" {
		fmt.Println("toBranch should be provided")
	}
	var gitHelper = git.GitHelper{FromBranch: fromBranch, ToBranch: toBranch}
	err := gitHelper.CheckOut()
	if err != nil {
		fmt.Println(err.Error())
	}
	flag.Parse()
}
