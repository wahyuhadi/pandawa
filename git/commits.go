package git

import (
	"context"
	"fmt"
	"net/http"
	"pandawa/github"
)

// GetListCommitsSha function to get sha for all commits in user repos
// example : GET https://api.github.com/repos/wahyuhadi/python-shodan-cli/commits?per_page=100000
func GetListCommitsSha(session *http.Client, user string, repos []string, perpage string) {
	client := github.NewClient(session)
	ctx := context.Background()

	for _, reponame := range repos {
		// ListCommitsSha function in github/repos_commits.go
		commits, _, err := client.Repositories.ListCommitsSha(ctx, user, reponame, perpage, nil)

		if err != nil {
			fmt.Println(err)
		}

		for _, sha := range commits {
			GetChangesCommit(session, user, reponame, *sha.SHA)
		}
	}
}

func GetChangesCommit(session *http.Client, user, reponame, sha string) {
	client := github.NewClient(session)
	ctx := context.Background()

	// GetComit function github/repos_commits.go
	commits, _, err := client.Repositories.GetCommit(ctx, user, reponame, sha)

	if err != nil {
		fmt.Println(err)
	}

	for _, changeCommits := range commits.Files {

		// Posible object patch in json response is nil
		// make validate to patch the error nil pointer
		if changeCommits.Patch != nil {

			// Change commit file changes with regex, return bool type is true or false
			// function RegexCheckCommit in regex.go file
			_, find, _ := RegexCheckCommit(*changeCommits.Patch)

			if find {
				fmt.Println("\n[+] repo name : ", reponame)
				fmt.Println("[+] Filename : ", *changeCommits.Filename)
				fmt.Println("[+] Add : ", *changeCommits.Additions)
				fmt.Println("[+] Deleted : ", *changeCommits.Deletions)
			}
		}
	}
}
