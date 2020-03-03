package git

import (
	"fmt"
	"net/http"
	"pandawa/github"
)

// Get function List  member from organtation
func GetListMemberFromOrg(session *http.Client, org string) ([]string, error) {
	members, _, err := github.LisMembers(session, org, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return members, nil
}
