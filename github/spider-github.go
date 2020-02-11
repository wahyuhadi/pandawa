package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Get User Info from API
type UsersInfo struct {
	Login             string      `json:"login"`
	ID                int         `json:"id"`
	NodeID            string      `json:"node_id"`
	AvatarURL         string      `json:"avatar_url"`
	GravatarID        string      `json:"gravatar_id"`
	URL               string      `json:"url"`
	HTMLURL           string      `json:"html_url"`
	FollowersURL      string      `json:"followers_url"`
	FollowingURL      string      `json:"following_url"`
	GistsURL          string      `json:"gists_url"`
	StarredURL        string      `json:"starred_url"`
	SubscriptionsURL  string      `json:"subscriptions_url"`
	OrganizationsURL  string      `json:"organizations_url"`
	ReposURL          string      `json:"repos_url"`
	EventsURL         string      `json:"events_url"`
	ReceivedEventsURL string      `json:"received_events_url"`
	Type              string      `json:"type"`
	SiteAdmin         bool        `json:"site_admin"`
	Name              string      `json:"name"`
	Company           string      `json:"company"`
	Blog              string      `json:"blog"`
	Location          string      `json:"location"`
	Email             interface{} `json:"email"`
	Hireable          bool        `json:"hireable"`
	Bio               string      `json:"bio"`
	PublicRepos       int         `json:"public_repos"`
	PublicGists       int         `json:"public_gists"`
	Followers         int         `json:"followers"`
	Following         int         `json:"following"`
	CreatedAt         time.Time   `json:"created_at"`
	UpdatedAt         time.Time   `json:"updated_at"`
}

var (
	URI        = "https://api.github.com/users/"
	PublicRepo = 0
	GithubUser = ""
)

// Client for connect to github API
var myClient = &http.Client{Timeout: 10 * time.Second}

func GithubReq(url string) (*http.Response, error) {
	res, err := myClient.Get(url)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func InitalSpider(login string) {
	// define username github
	GithubUser = login
	isGetUserInfoUri := URI + login
	r, err := GithubReq(isGetUserInfoUri)
	if err != nil {
		fmt.Println("[!] Error when access user data from github")
		fmt.Println("[!] Cek this endpoint ", URI)
		return
	}

	defer r.Body.Close()
	// new decoder stuff
	userInfoRes := new(UsersInfo)
	// Decode
	json.NewDecoder(r.Body).Decode(userInfoRes)
	// Parsing
	PublicRepo = userInfoRes.PublicRepos

	fmt.Println("[+] Spider github account :", GithubUser)
	fmt.Println("[+] Work at : ", userInfoRes.Company)
	fmt.Println("[+] Total public repo : ", PublicRepo)
}
