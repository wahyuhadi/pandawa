package git

import (
	"context"
	"fmt"
	"os"

	"golang.org/x/oauth2"
)

func InitialSearch(user string) {
	gitoken, err := GetInitialToken()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ctx := context.Background()

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: gittoken},
	)

	tc := oauth2.NewClient(ctx, ts)

	repos, _ := GetList(tc, user, "100000")
	GetListCommitsSha(tc, user, repos, "100000")

}

// GetUserFromOrg to get list user from org
// and save from db ->
func GetUserFromOrg(org string) {
	gitoken, err := GetInitialToken()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: gittoken},
	)

	tc := oauth2.NewClient(ctx, ts)

}
