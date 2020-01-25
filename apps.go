package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	db "pandawa/database"
	"pandawa/github"
	shodan "pandawa/shodan"
)

var (
	domain = flag.String("domain", "google.com", "Domain for osint")
	config = flag.String("config", "~/.pandawa-config.json", "Specify the configuration file.")

	operation = flag.String("ops", "aurora", "Operation name")
)

// Membaca file configurasi dari config.json
type Configuration struct {
	Shodan struct {
		Key string
	}
	Listen struct {
		Address string
		Port    string
	}
	OutboundCall struct {
		CallerID  string
		Retries   int
		SpoolPath string
	}
	VmRoot string
}

//  Membaca file conf untuk mendapatkan key / API KEY shodan dan dll
func ReadFileConf() Configuration {
	flag.Parse()
	file, err := os.Open(*config)
	if err != nil {
		log.Fatal("can't open config file: ", err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	Config := Configuration{}
	err = decoder.Decode(&Config)
	if err != nil {
		log.Fatal("can't decode config JSON: ", err)
	}
	return Config

}

func main() {
	fmt.Println("[+] Pandawa osint")
	flag.Parse()
	fmt.Println(*domain)

	dbname := db.GenerateDb(*operation)
	fmt.Println("[+] Location DB ", dbname)

	file := ReadFileConf()

	// shodan search
	// Mencari dengan mengunakan shodan mmh3
	// req : shodan key
	shodan.PreSearch(file.Shodan.Key)

	// search keyword didalam code github
	// req : keyword and order type (asc , desc)
	db.GenDbGithub(dbname)
	github.GetGitRepo(*domain, "desc")

}
