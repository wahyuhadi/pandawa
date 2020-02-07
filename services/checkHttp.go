package services

import (
	"fmt"
	"net/http"

	httpreq "github.com/wahyuhadi/httpreq"
)

//  Check http connections in ip address
func CheckHttpService(host string) {
	file := [4]string{".git/config", ".env", ".env-example", ".env-production"}

	open := CheckPort(host, "80")
	if open {
		for i := 0; i < len(file); i++ {
			//fmt.Println(host + file[i])
			isHost := "http://" + host + "/" + file[i]
			GetRequests(isHost)
		}
	}
}

//  Check https connections in ip address
func CheckHttpsService(host string) {
	file := [4]string{".git/config", ".env", ".env-example", ".env-production"}

	open := CheckPort(host, "443")
	if open {
		for i := 0; i < len(file); i++ {
			//fmt.Println(host + file[i])
			isHost := "https://" + host + "/" + file[i]
			GetRequests(isHost)
		}
	}
}

// get requests
func GetRequests(host string) {
	r, err := httpreq.Req(host)
	if err != nil {
		//fmt.Println(err)
		return
	}

	if r.StatusCode == http.StatusOK {
		fmt.Println("  ----------------------> ", host)
	}
}
