package shodan

import (
	"bufio"
	"encoding/base64"
	"io/ioutil"
	"os"

	mmh3 "pandawa/murmur3"
)

// GetFavicon
// get Favicon dari URL untuk digunakan search osint di shodan dengan MMH3
func ToB64(nameico string) uint32 {
	f, _ := os.Open(nameico)

	// Read entire JPG into byte slice.
	reader := bufio.NewReader(f)
	content, _ := ioutil.ReadAll(reader)

	// Encode as base64.
	encoded := base64.StdEncoding.EncodeToString(content)
	// Print encoded data to console.
	// ... The base64 image can be used as a data URI in a browser.
	return FaviconToMMH(encoded)
}

// Melakuka hash dengan menggunakan mmh3
// mmh3(base64(favicon.ico))
func FaviconToMMH(b64enc string) uint32 {
	mmh3_64 := mmh3.Sum32WithSeed([]byte(b64enc), 64)
	return mmh3_64
}
