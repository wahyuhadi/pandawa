package git

import (
	"context"
	"fmt"
	"os"
	"pandawa/database"

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
		&oauth2.Token{AccessToken: gitoken},
	)

	tc := oauth2.NewClient(ctx, ts)

	repos, _ := GetList(tc, user, "100000")
	GetListCommitsSha(tc, user, repos, "100000")

}

// GetUserFromOrg to get list user from org
// and save from db ->
func GetUserFromOrg(org, ops string) {
	gitoken, err := GetInitialToken()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: gitoken},
	)
	tc := oauth2.NewClient(ctx, ts)

	// generate tables if ops not null
	// with name table : github_org
	if ops != "" {
		query := "CREATE TABLE IF NOT EXISTS github_org (id INTEGER PRIMARY KEY, user TEXT, url TEXT, org TEXT)"
		operation := "pandawa-output/" + ops + ".db"
		database.GenerateTables(operation, query)
	}

	members, err := GetListMemberFromOrg(tc, org)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, member := range members {
		fmt.Println("[+] User   : ", *member.Login)
		fmt.Println("[+] Github : ", *member.HTMLURL, "\n")

		operation := "pandawa-output/" + ops + ".db"
		// queryInsert := "INSERT INTO github_org (" + *member.Login + "," + *member.HTMLURL + "," + org + ") VALUES (?, ?, ?)"
		database.InsertDataUserOrg(operation, *member.Login, *member.HTMLURL, org)
	}

	return
	// fmt.Println(members)

}
