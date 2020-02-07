package database

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func GenDbShodan(nameDB string) {
	database, _ := sql.Open("sqlite3", nameDB)
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS shodan (id INTEGER PRIMARY KEY, ip INTEGER, ip_str TEXT,  isp TEXT)")
	statement.Exec()
}

func AddDBShodan(ip int, ip_str, isp, nameDB string) {
	GenDbShodan(nameDB)
	database, _ := sql.Open("sqlite3", nameDB)
	statement, _ := database.Prepare("INSERT INTO shodan (ip, ip_str, isp) VALUES (?, ?, ?)")
	statement.Exec(ip, ip_str, isp)
}

func CollectDataShodan(nameDB string) {
	dblocation := "pandawa-output/" + nameDB + ".db"
	fmt.Println("[+] ID : ", dblocation)
	database, _ := sql.Open("sqlite3", dblocation)
	rows, _ := database.Query("SELECT * FROM shodan")
	var id int
	var ip int
	var ip_str string
	var isp string
	for rows.Next() {
		rows.Scan(&id, &ip, &ip_str, &isp)
		fmt.Println("*************************************************")
		fmt.Println("[+] ID 	   :", strconv.Itoa(id))
		fmt.Println("[+] IP 	   :", ip)
		fmt.Println("[+] IP_STR :", ip_str)
		fmt.Println("[+] ISP    :", isp)

	}
}
