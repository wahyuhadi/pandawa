package spider

import (
	"fmt"

	"golang.org/x/net/html"
)

func InitJs(URI string) {
	page, err := Parse(URI)
	if err != nil {
		fmt.Printf("Error getting page %s %s\n", URI, err)
		return
	}

	linksJs := pageJs(nil, page)
	for _, js := range linksJs {
		fmt.Println("[js] ---> ", js)
	}
}

func pageJs(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "script" {
		for _, a := range n.Attr {
			if a.Key == "src" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = pageJs(links, c)
	}
	return links
}
