package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/ravishankarsrrav/git-cherry-picker/pkg/git"
	"github.com/ravishankarsrrav/git-cherry-picker/pkg/ui"
	"github.com/ravishankarsrrav/git-cherry-picker/pkg/utils"
	flag "github.com/spf13/pflag"
)

var fromBranch string
var toBranch string
var continueCP bool

func main() {
	flag.StringVar(&fromBranch, "from_branch", "", "Name of the branch from which the commit is cherry picked")
	flag.StringVar(&toBranch, "to_branch", "", "Name of the branch to which the commit is cherry picked. (Use this only when you are on different branch)")
	flag.BoolVar(&continueCP, "continue", false, "Flag to continue cherry pick after the merge conflict is resolved")
	flag.Parse()
	color.HiGreen("gcpk tool lets you easily cherry-pick the commit with an easy terminal interface 🙌")
	color.HiBlue("Usage \n")
	color.HiBlue("gcpk [flags]")
	color.HiBlue("Flags:")
	color.HiBlue("       --from_branch   Name of the branch from which the commit is cherry picked")
	color.HiBlue("       --to_branch     Name of the branch to which the commit is cherry picked. (Use this only when you are on different branch)")
	color.HiBlue("       --continue      Flag to continue cherry pick after the merge conflict is resolved \n")
	color.HiBlue("Use 'gcpk --help' for more information about the flags")
	color.HiBlue("Support: If you have any questions, file an issue at https://github.com/ravishankarsrrav/git-cherry-picker/issues/new")
	color.HiBlue(": If you are using it, star a github repo at https://github.com/ravishankarsrrav/git-cherry-picker \n\n\n")
	var gitHelper = git.GitHelper{FromBranch: fromBranch, ToBranch: toBranch}
	if continueCP {
		err := gitHelper.AddAllChanges()
		if err != nil {
			color.Red("Unable to all the changes.")
		}
		err = gitHelper.Continue()
		if err != nil {
			color.Red("Unable to continue the cherry-pick process.")
			return
		}
		color.HiYellow("Cherry pick process is completed!")
		return
	}

	if fromBranch == "" {
		color.Red("from_branch flag is required.")
		color.HiGreen("Please enter the branch name from which the commit will be cherry picked?")
		_, err := fmt.Scanf("%s", &fromBranch)
		if err != nil {
			color.Red(err.Error())
		}
		if fromBranch == "" {
			color.Red("You cannot proceed without adding the from_branch")
			return
		}
	}

	gitHelper = git.GitHelper{FromBranch: fromBranch, ToBranch: toBranch}

	if toBranch != "" {
		// checkout to the branch
		err := gitHelper.CheckOut()
		if err != nil {
			color.Red("checkout failed %s\n", err.Error())
		}
	}

	// fetch all the commits
	commits, _ := gitHelper.ListCommits()
	// check if there are no commits in the branch
	if len(commits) <= 0 {
		color.Red("There are no commits in the %s branch.\n", fromBranch)
		return
	}
	// render the list of the commits
	selectedElement := ui.Draw(commits)

	// Print help message if the commit message is unselected
	if selectedElement == 0 {
		color.Red("Commit message should be selected from the list to process cherry-pick.")
		return
	}

	// extract the commit id
	commitID := utils.ExtractCommitId(commits[selectedElement-1])

	// cherry-pick the commit
	err := gitHelper.CherryPick(commitID)
	if err != nil {
		color.Red("Cherry Pick Failed!")
		return
	}
	color.HiGreen("Cherry pick is done!")
	return
}
