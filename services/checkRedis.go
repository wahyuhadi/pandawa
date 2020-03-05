package services

func CheckRedis(host string) {
	open := CheckPort(host, "6379")
	if open {
		fmt.Println("     --------> Redis ", host, " with port 6379")
	}
}
