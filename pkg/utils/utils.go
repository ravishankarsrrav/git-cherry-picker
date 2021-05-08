package utils

import "strings"

/*ExtractCommitId extracts the commit hash from the commit message*/
func ExtractCommitId(commitMessage string) string {
	splittedString := strings.Split(commitMessage, " ")
	return splittedString[0]
}
