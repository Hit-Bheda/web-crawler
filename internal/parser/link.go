package parser

import (
	"fmt"

	"golang.org/x/net/html"
)

func LinkParser(n *html.Node) {

	if n.Type == html.ElementNode {
		// fmt.Println("Node :", n)
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				fmt.Println(attr.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		LinkParser(c)
	}
}
