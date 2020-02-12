package github

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

// Client for connect to github API
var myClient = &http.Client{Timeout: 10 * time.Second}

// GithubReq request get github repo
func GithubReq(url string) (*http.Response, error) {
	githubToken := os.Getenv("shodan")

	if githubToken == "" {
		fmt.Println("[!] github token not found in yout environtment please add export github=xxxxxxx")
		os.Exit(1)
	}

	// isToken := "token " + githubToken
	res, err := myClient.Get(url)
	res.Header.Set("Authorization", "token 49fc0374d5ed3ac4febd85b327f49721de593f83")
	fmt.Println(res.Header)
	if err != nil {
		return nil, err
	}
	return res, nil
}
