package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)

var (
	PerPages   = 100000
	URIShaRepo = "https://api.github.com/repos/" //"https://api.github.com/repos/wahyuhadi/GitTools/commits"
)

type ShaCommit []struct {
	Sha     string `json:"sha"`
	Parents []struct {
		Sha     string `json:"sha"`
		URL     string `json:"url"`
		HTMLURL string `json:"html_url"`
	} `json:"parents"`
}

// Getshacommiter
func GetShaCommiter(repoName string) {
	isURIsha := URIShaRepo + GithubUser + "/" + repoName + "/commits?per_page=" + strconv.Itoa(PerPages)
	r, err := GithubReq(isURIsha)
	if err != nil {
		fmt.Println("[!] Error when access user data from github")
		fmt.Println("[!] Cek this endpoint ", isURIsha)
		return
	}

	// body save
	body, err := ioutil.ReadAll(r.Body)
	r.Body.Close()
	if err != nil {
		fmt.Println("[!] Error func GetShaCommiter  ", err)
		return
	}

	// Parsing
	data := ShaCommit{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("[!] Error func GetShaCommiter ")
		return
	}

	for _, x := range data {
		fmt.Println("-------> [+] Check sha commit ", x.Sha)
	}
}
