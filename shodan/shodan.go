package shodan

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	db "pandawa/database"
)

// Validasi semua inputan sudah terpenui atau tidak
// Jika terpenuhi makan akan melakukan pencarian jika tidak makan osint pada shodan tidak dilanjutkan
//

var (
	URI = "https://api.shodan.io/shodan/host/search?key="
)

// ShodanResponses for response API model
// api.shodan.io -> Documentation
type ShodanResponses struct {
	Matches []struct {
		IP    int    `json:"ip"`
		IPStr string `json:"ip_str"`
		Isp   string `json:"isp"`
	} `json:"matches"`
}

func PreSearch(key string, mmh3 uint32, dbname string) {
	fmt.Println(mmh3)

	// for temp search by query like this
	// before mumur function search done
	// search by http.hash
	search := "tiket.com"
	isQuery := "&query=" + search
	isURI := URI + key + isQuery

	// Initial request to URI
	res, err := http.Get(isURI)
	if err != nil {
		log.Fatal(err)
	}

	// Read response body from request
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	// check status code
	if res.StatusCode != http.StatusOK {
		log.Fatal("Unexpected status code", res.StatusCode)
	}

	// Marshall json response from Struct
	data := ShodanResponses{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal(err)
	}

	// Add data to shodan tables
	AddDataShodanToDB(dbname, data)
}

func AddDataShodanToDB(dbname string, data ShodanResponses) {
	for _, i := range data.Matches {
		db.AddDBShodan(i.IP, i.IPStr, i.Isp, dbname)
	}
}
