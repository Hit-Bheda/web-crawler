package parser

import (
	"os"
	"strings"

	"golang.org/x/net/html"
)

func TextParser(n *html.Node, file *os.File) {

	if n.Type == html.TextNode && n.Parent != nil {
		if n.Parent.Data != "script" && n.Parent.Data != "style" {
			text := strings.TrimSpace(n.Data)
			if text != "" {
				file.WriteString(text + "\n")
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		TextParser(c, file)
	}
}
