package services

func PreInit(host string) {
	// check http connection
	CheckHttpService(host)
	// check https connection
	CheckHttpsService(host)
	// check mongodb with port 27017
	CheckMongo_27017(host)
	// check mongodb with port 27018
	CheckMongo_27018(host)
}
