package github

import (
	"net/http"
	"time"
)

// Client for connect to github API
var myClient = &http.Client{Timeout: 10 * time.Second}

// GithubReq request get github repo
func GithubReq(url string) (*http.Response, error) {
	res, err := myClient.Get(url)
	if err != nil {
		return nil, err
	}
	return res, nil
}
