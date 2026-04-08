package parser

import (
	"errors"
	"strings"
	"unicode"

	"golang.org/x/net/html"
)

var noiseTags = map[string]bool{
	"script":  true,
	"style":   true,
	"nav":     true,
	"header":  true,
	"footer":  true,
	"aside":   true,
	"meta":    true,
	"link":    true,
	"noscript": true,
	"input":   true,
	"button":  true,
	"select":  true,
	"textarea": true,
	"form":    true,
	"iframe":  true,
	"canvas":  true,
	"svg":     true,
	"path":    true,
}

var skipParents = map[string]bool{
	"nav":     true,
	"header":  true,
	"footer":  true,
	"aside":   true,
	"menu":    true,
	"sidebar": true,
	"advertisement": true,
	"ad":      true,
}

func isNoiseTag(n *html.Node) bool {
	if n.Type != html.ElementNode {
		return false
	}
	return noiseTags[strings.ToLower(n.Data)]
}

func hasNoiseParent(n *html.Node) bool {
	for p := n.Parent; p != nil; p = p.Parent {
		if p.Type == html.ElementNode && skipParents[strings.ToLower(p.Data)] {
			return true
		}
	}
	return false
}

func TextParser(n *html.Node) string {
	builder := strings.Builder{}
	text := extractText(n)
	cleaned := cleanText(text)
	builder.WriteString(cleaned)
	return builder.String()
}

func extractText(n *html.Node) string {
	builder := strings.Builder{}
	
	if n.Type == html.TextNode && n.Parent != nil {
		if !isNoiseTag(n.Parent) && !hasNoiseParent(n.Parent) {
			text := n.Data
			if text != "" {
				builder.WriteString(text)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		childData := extractText(c)
		builder.WriteString(childData)
	}
	return builder.String()
}

func cleanText(text string) string {
	var result strings.Builder
	prevSpace := false
	
	for _, r := range text {
		if unicode.IsSpace(r) {
			if !prevSpace {
				result.WriteRune(' ')
				prevSpace = true
			}
		} else {
			result.WriteRune(r)
			prevSpace = false
		}
	}
	
	cleaned := strings.TrimSpace(result.String())
	cleaned = strings.ReplaceAll(cleaned, " ,", ",")
	cleaned = strings.ReplaceAll(cleaned, " .", ".")
	
	return cleaned
}

func GetTitle(n *html.Node) (string, error) {
	builder := strings.Builder{}
	if n.Type == html.ElementNode && strings.ToLower(n.Data) == "title" {
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
