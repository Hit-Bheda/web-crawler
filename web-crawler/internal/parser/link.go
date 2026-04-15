package parser

import (
	"net/url"

	"github.com/rs/zerolog"
	"golang.org/x/net/html"
)

func LinkParser(n *html.Node, pageUrl string, log zerolog.Logger) ([]string, error) {
	var links []string
	if n.Type == html.ElementNode {
		baseUrl, err := url.Parse(pageUrl)
		if err != nil {
			log.Error().Err(err).Str("url", pageUrl).Msg("Failed to get base url")
			return nil, err
		}

		for _, attr := range n.Attr {
			if attr.Key == "href" && attr.Val != "" {
				href, err := url.Parse(attr.Val)
				if err != nil {
					continue
				}

				resolved := baseUrl.ResolveReference(href)
				resolved.Fragment = ""
				links = append(links, resolved.String())
			}
		}

	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		childLinks, _ := LinkParser(c, pageUrl, log)
		links = append(links, childLinks...)
	}

	return links, nil
}
