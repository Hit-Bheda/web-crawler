package parser

import (
	"errors"
	"strings"

	"github.com/rs/zerolog"
	"golang.org/x/net/html"
)

func LinkParser(n *html.Node, url string, log zerolog.Logger) ([]string, error) {
	var links []string
	if n.Type == html.ElementNode {
		baseUrl, err := GetBaseUrl(url)
		if err != nil {

			log.Error().Err(err).Str("url", url).Msg("Failed to get base url")
			return nil, err
		}

		for _, attr := range n.Attr {
			if attr.Key == "href" {
				rawUrl := strings.Split(attr.Val, "/")
				if rawUrl[0] == "http:" || rawUrl[0] == "https:" {
					links = append(links, string(attr.Val))
				} else if attr.Val[0] == '/' {
					links = append(links, string(baseUrl+string(attr.Val)))
				}
			}
		}

	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		childLinks, _ := LinkParser(c, url, log)
		links = append(links, childLinks...)
	}

	return links, nil
}

func GetBaseUrl(givenUrl string) (string, error) {
	sptUrl := strings.Split(givenUrl, "/")
	baseUrl := ""
	if sptUrl[0] == "http:" || sptUrl[0] == "https:" {
		baseUrl = sptUrl[0] + "//" + sptUrl[2]
	} else {
		return "", errors.New("Invalid base url!")
	}
	return baseUrl, nil
}
