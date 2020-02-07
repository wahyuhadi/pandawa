package shodan

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Validasi semua inputan sudah terpenui atau tidak
// Jika terpenuhi makan akan melakukan pencarian jika tidak makan osint pada shodan tidak dilanjutkan
//

var (
	URI = "https://api.shodan.io/shodan/host/search?key="
)

type ShodanResponses []struct {
	IP    int    `json:"ip"`
	IPStr string `json:"ip_str"`
	Isp   string `json:"isp"`
}

func PreSearch(key string, mmh3 uint32, dbname string) {
	fmt.Println(key, mmh3)
	search := "tiket.com"
	isQuery := "&query=" + search
	isURI := URI + key + isQuery

	res, err := http.Get(isURI)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		log.Fatal("Unexpected status code", res.StatusCode)
	}
	data := ShodanResponses{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
	for _, i := range data {
		fmt.Println(i.IPStr)
	}
	//AddDataShodanToDB(dbname, data)
}

func AddDataShodanToDB(dbname string, data ShodanResponses) {

}
