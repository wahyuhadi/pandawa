package shodan

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var (
	fileName    string
	fullUrlFile string
)

func GetFav(name string) uint32 {

	fullUrlFile = name

	// Build fileName from fullPath
	filename := buildFileName()

	// Create blank file
	file := createFile()

	// Put content on file
	putFile(file, httpClient())
	return ToB64(filename)

}

func putFile(file *os.File, client *http.Client) {
	resp, err := client.Get(fullUrlFile)

	checkError(err)

	defer resp.Body.Close()

	size, err := io.Copy(file, resp.Body)

	defer file.Close()

	checkError(err)

	fmt.Println("[+] Download favicon file %s with size %d", fileName, size)
}

func buildFileName() string {
	fileUrl, err := url.Parse(fullUrlFile)
	checkError(err)

	path := fileUrl.Path
	segments := strings.Split(path, "/")

	fileName = segments[len(segments)-1]
	return fileName
}

func httpClient() *http.Client {
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	return &client
}

func createFile() *os.File {
	file, err := os.Create(fileName)

	checkError(err)
	return file
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
