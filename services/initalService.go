package services

func PreInit(host string) {
	CheckHttpService(host)
	CheckHttpsService(host)
}
