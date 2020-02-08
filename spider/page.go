package spider

import (
	"fmt"

	"golang.org/x/net/html"
)

func InitPage(URI string) {
	page, err := Parse(URI)
	if err != nil {
		fmt.Printf("Error getting page %s %s\n", URI, err)
		return
	}

	links := pageLinks(nil, page)
	for _, link := range links {
		fmt.Println("[page] ---> ", link)
	}
}

func pageLinks(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = pageLinks(links, c)
	}
	return links
}
