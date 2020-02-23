package git

import (
	"context"
	"fmt"
	"os"

	"golang.org/x/oauth2"
)

func InitialSearch(user string) {
	gittoken := os.Getenv("github")

	if gittoken == "" {
		fmt.Println("[!] github token not found in env / .bashrc / .zshrc")
		os.Exit(1)
		return
	}

	ctx := context.Background()

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: gittoken},
	)

	tc := oauth2.NewClient(ctx, ts)

	repos, _ := GetList(tc, user, "100000")
	GetListCommitsSha(tc, user, repos, "100000")

}
