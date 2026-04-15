package fetcher

import (
	"github.com/rs/zerolog"
	"golang.org/x/net/html"
	"net/http"
)

func FetchDocument(url string, log zerolog.Logger) (*html.Node, error) {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 ...")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Error().Err(err).Str("url", url).Msg("Failed to fetch page data")
		return nil, err
	}

	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Error().Err(err).Str("url", url).Msg("Failed to fetch page data")
		return nil, err
	}
	return doc, nil
}
