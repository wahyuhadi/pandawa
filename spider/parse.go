package spider

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

func Parse(url string) (*html.Node, error) {
	r, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Cannot get page")
	}

	b, err := html.Parse(r.Body)
	if err != nil {
		return nil, fmt.Errorf("Cannot get page")
	}

	return b, err
}
