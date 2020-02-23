package git

import (
	"context"
	"errors"
	"net/http"
	"pandawa/github"
)

// GetList Repo function -> to get all list repo in user
// --> users/userlogin/repos?per_page=10000
func GetList(session *http.Client, login, perpage string) ([]string, error) {
	client := github.NewClient(session)
	ctx := context.Background()

	opt := &github.RepositoryListOptions{
		Type:    "sources",
		Perpage: "100000",
	}

	repos, _, err := client.Repositories.ListName(ctx, login, perpage, opt)
	if err != nil {
		return nil, errors.New("[!] error when get list repo")
	}

	var RepositoryName []string
	for _, repo := range repos {
		// get repo only fork == false
		if *repo.Fork == false {
			RepositoryName = append(RepositoryName, *repo.Name)
		}
	}

	// Return array name
	return RepositoryName, nil
}
