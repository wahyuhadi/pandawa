package git

import (
	"context"
	"fmt"
	"net/http"
	"pandawa/github"
)

// Get function List  member from organtation
func GetListMemberFromOrg(session *http.Client, org string) ([]*github.User, error) {

	client := github.NewClient(session)
	ctx := context.Background()

	members, _, err := client.Organizations.ListMembers(ctx, org, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return members, nil
}
