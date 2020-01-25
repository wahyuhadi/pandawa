package database

import (
	"fmt"
	"log"
	"os"
)

func GenerateDb(operation string) string {
	_ = os.Mkdir("pandawa-output", os.ModePerm)
	name := "pandawa-output/" + operation + ".db"

	fmt.Println("[+] Create database for operation ", operation, "....")
	_, err := os.Create(name)
	if err != nil {
		log.Fatal("[!] Error when creating db .", err)
		os.Exit(1)
	}
	return name
}
