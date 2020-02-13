package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type ErrorMessage struct {
	Message          string `json:"message"`
	DocumentationURL string `json:"documentation_url"`
}

// Client for connect to github API
var myClient = &http.Client{Timeout: 10 * time.Second}

// GithubReq request get github repo
func GithubReq(url string) (*http.Response, error) {
	githubToken := os.Getenv("github")

	if githubToken == "" {
		fmt.Println("[!] github token not found in yout environtment please add export github=xxxxxxx")
		os.Exit(1)
	}

	//isToken := "token " + os.Getenv("github")
	res, err := myClient.Get(url)
	//res.Header.Add("Authorization", isToken)
	//res.Header.Add("Accept", "application/vnd.github.v3.text-match+json")
	//fmt.Println(res.Header)

	var xRateLimitRemaining string = res.Header["X-Ratelimit-Remaining"][0]
	fmt.Println("[x] X rate limit ", xRateLimitRemaining)
	if res.StatusCode == 403 {
		body, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			fmt.Println("[!] Error parsing 403 response ")
			//return nil, err
		}

		// Parsing
		data := ErrorMessage{}
		err = json.Unmarshal(body, &data)
		if err != nil {
			fmt.Println("[!] Error func Prespider when marshalling ")
			//return nil, err
		}

		fmt.Println("[!] Error Message : ", data.Message)
		fmt.Println("[!] Document : ", data.DocumentationURL)
		//return nil, err
	}

	if err != nil {
		return nil, err
	}
	return res, nil
}
