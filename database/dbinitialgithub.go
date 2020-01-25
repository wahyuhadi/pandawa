package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// Generate DB  untuk github
//
func GenDbGithub(nameDB string) {
	database, _ := sql.Open("sqlite3", nameDB)
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS github (id INTEGER PRIMARY KEY, repository TEXT, updateat TEXT, rawurl TEXT, giturl TEXT)")
	statement.Exec()
}

// Insert function where from database
// insert
func AddGithubData(repository, updateat, rawurl, giturl, nameDB string) {
	database, _ := sql.Open("sqlite3", nameDB)
	statement, _ := database.Prepare("INSERT INTO github (repository, updateat,rawurl, giturl) VALUES (?, ?, ?, ?)")
	statement.Exec(repository, updateat, rawurl, giturl)
}

// Collect data from db
func CollectData(nameDB string) {

}
