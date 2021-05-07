package main

import (
	flag "github.com/spf13/pflag"
)

var fromBranch string
var toBranch string

func main() {
	flag.StringVar(&fromBranch, "from_branch", "", "name of the branch from which the commit is cherry picked")
	flag.StringVar(&toBranch, "to_branch", "", "name of the branch to which the commit is cherry picked")
	flag.Parse()
}
