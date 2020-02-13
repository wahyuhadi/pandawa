package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
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
	PerPage    = 100000
)

// InitialSpider function for get user information
// - public repo - company name - etc
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
	PreSpider()
}

type RepoPublic []struct {
	ID     int    `json:"id"`
	NodeID string `json:"node_id"`
	Name   string `json:"name"`
	Fork   bool   `json:"fork"`
}

// Get hash from repo public list
func PreSpider() {
	// isPage := CalculatePage()
	fmt.Println("[+] This process starting, take some coffee and enjoyed .. ")
	URIRepoPub := URI + GithubUser + "/repos?per_page=" + strconv.Itoa(PerPage)
	isFinalRepo := URIRepoPub

	// Get
	r, err := GithubReq(isFinalRepo)
	if err != nil {
		fmt.Println("[!] Error when access user data from github")
		fmt.Println("[!] Cek this endpoint ", URI)
		return
	}

	// body save
	body, err := ioutil.ReadAll(r.Body)
	r.Body.Close()
	if err != nil {
		fmt.Println("[Debug] Error func Prespider when body save ")
		return
	}

	// Parsing
	data := RepoPublic{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("[Debug] Error func Prespider when marshalling Parsing ")
		return
	}

	// Loop
	loop := 0
	for _, x := range data {
		if x.Fork == false {
			loop = loop + 1
			fmt.Println("[", loop, "] pre-cloning repo : ", x.Name)
			GetShaCommiter(x.Name)
		}
	}
}
