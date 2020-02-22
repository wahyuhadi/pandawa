package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	db "pandawa/database"

	//"pandawa/shodan"

	prompt "pandawa/cli"
	shodan "pandawa/shodan"
)

var (
	domain    = flag.String("domain", "google.com", "Domain for osint")
	config    = flag.String("config", "~/.pandawa-config.json", "Specify the configuration file.")
	favicon   = flag.String("fav", "no", "favicon.ico URL ")
	operation = flag.String("ops", "no", "Operation name")
	cli       = flag.String("cli", "no", "cli mode connect to sqldatabase")
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

	// pada saat mode cli digunakan untuk melakukan query pada database sqlite
	// next nya akan dikembangkan untuk terminal command
	if *cli == "yes" {
		fmt.Println("[+] Enter operation name", *operation)
		prompt.Cli()
		os.Exit(1)
	}

	if *operation == "no" {
		fmt.Println("[+] Enter operation name example --ops=pandawa")
		os.Exit(1)
	}

	if *favicon == "no" {
		fmt.Println("[+] Enter favicon url  example --fav=https://localhost.com/favicon.ico")
		os.Exit(1)
	}

	dbname := db.GenerateDb(*operation)
	fmt.Println("[+] Location DB ", dbname)

	//file := ReadFileConf()
	shodankey := os.Getenv("shodan")
	if shodankey == "" {
		fmt.Println("[!] shodan key not found please export shodan key in .bashrc or .zshrc")
		os.Exit(1)
	}
	// convert favicon ico ke mumurhash
	// mmh3 := shodan.GetFav(*favicon)
	mmh3 := "tes"
	//Exec shodan main
	shodan.PreSearch(shodankey, mmh3, dbname)

	// search keyword didalam code github
	// req : keyword and order type (asc , desc)
	// db.GenDbGithub(dbname)
	// github.GetGitRepo(*domain, "desc", dbname)

}
