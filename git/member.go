package git

import (
	"context"
	"errors"
	"net/http"
	"pandawa/github"
)

// Get function List  member from organtation
func GetListMemberFromOrg(session *http.Client, org string) ([]*github.User, error) {

	client := github.NewClient(session)
	ctx := context.Background()

	// ListsMembers in file github/orgs_members.go file
	members, _, err := client.Organizations.ListMembers(ctx, org, nil)
	if err != nil {
		return nil, errors.New("Error when get users from github org")
	}

	return members, nil
}
