package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// Generate DB  untuk github
//
func GenDbGithub(nameDB string) {
	database, _ := sql.Open("sqlite3", nameDB)
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS github (id INTEGER PRIMARY KEY, repository TEXT, updateat TEXT, rawurl TEXT)")
	statement.Exec()
}
