package services

import "fmt"

func CheckMongo_27017(host string) {
	open := CheckPort(host, "27017")
	if open {
		fmt.Println("     --------> mongodb ", host, " with port 27017")
	}
}

func CheckMongo_27018(host string) {
	open := CheckPort(host, "27017")
	if open {
		fmt.Println("     --------> mongodb ", host, " with port 27018")
	}
}

// func TryToConnectMongo() {

// }
