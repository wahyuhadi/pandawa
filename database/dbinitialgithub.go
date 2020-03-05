package database

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

// Generate DB  untuk github
//
func GenDbGithub(nameDB string) {
	database, _ := sql.Open("sqlite3", nameDB)
	defer database.Close()
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS github (id INTEGER PRIMARY KEY, repository TEXT, updateat TEXT, rawurl TEXT, giturl TEXT)")
	statement.Exec()
}

// Insert function where from database
// insert
func AddGithubData(repository, updateat, rawurl, giturl, nameDB string) {
	database, _ := sql.Open("sqlite3", nameDB)
	defer database.Close()
	statement, _ := database.Prepare("INSERT INTO github (repository, updateat,rawurl, giturl) VALUES (?, ?, ?, ?)")
	statement.Exec(repository, updateat, rawurl, giturl)
}

// Collect data from db
func CollectData(nameDB string) {
	dblocation := "pandawa-output/" + nameDB + ".db"
	fmt.Println("[+] ID : ", dblocation)

	database, _ := sql.Open("sqlite3", dblocation)
	rows, _ := database.Query("SELECT * FROM github")
	if rows == nil {
		fmt.Println("[!] Data in github not found")
		return
	}

	defer database.Close()
	defer rows.Close()
	var id int
	var repository string
	var updateat string
	var rawurl string
	var giturl string
	for rows.Next() {
		rows.Scan(&id, &repository, &updateat, &rawurl, &giturl)
		fmt.Println("*************************************************")
		fmt.Println("[+] ID : ", strconv.Itoa(id))
		fmt.Println("[+] Repo :", repository)
		fmt.Println("[+] Update :", updateat)
		fmt.Println("[+] User : ", rawurl)
		fmt.Println("[+] URL : ", giturl)

	}
}

// Generate Tables function fror saved
func GenerateTables(nameDB, query string) {
	database, _ := sql.Open("sqlite3", nameDB)
	defer database.Close()
	statement, _ := database.Prepare(query)
	statement.Exec()
}

func CreateData(nameDB, query string) {
	// DB := nameDB + ".db"
	database, _ := sql.Open("sqlite3", nameDB)
	defer database.Close()
	statement, _ := database.Prepare(query)
	statement.Exec()
}
