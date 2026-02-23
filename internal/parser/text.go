package parser

import (
	"errors"
	"strings"

	"golang.org/x/net/html"
)

func TextParser(n *html.Node) string {
	builder := strings.Builder{}
	if n.Type == html.TextNode && n.Parent != nil {
		if n.Parent.Data != "script" && n.Parent.Data != "style" {
			text := strings.TrimSpace(n.Data)
			if text != "" {
				builder.WriteString(text)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		childData := TextParser(c)
		builder.WriteString(childData)
	}
	return builder.String()
}

func GetTitle(n *html.Node) (string, error) {
	builder := strings.Builder{}
	if n.Type == html.ElementNode && n.Data == "title" {
		if n.FirstChild != nil {
			title := strings.TrimSpace(n.FirstChild.Data)
			builder.WriteString(title)
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		childData, _ := GetTitle(c)
		builder.WriteString(childData)
	}

	if builder.String() == "" {
		return "", errors.New("Failed to get title!")
	}
	return builder.String(), nil
}
