package git

import (
	"fmt"
	"net/http"
	"pandawa/github"
)

func GetListMemberFromOrg(session *http.Client, org string) {
	members, _, err := github.ListMembers(session, org, nil)
	if err != nil {
		fmt.Println(err)

	}

}
