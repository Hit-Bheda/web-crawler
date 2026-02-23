package fetcher

import (
	// "encoding/json"
	// "fmt"
	"net/http"
	// "os"

	// "github.com/hit-bheda/web-crawler/internal/hash"
	// "github.com/hit-bheda/web-crawler/internal/parser"
	"github.com/rs/zerolog"
	"golang.org/x/net/html"
)

func FetchDocument(url string, log zerolog.Logger) *html.Node {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 ...")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Error().Err(err).Str("url", url).Msg("Failed to fetch page data")
	}

	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Error().Err(err).Str("url", url).Msg("Failed to fetch page data")
	}
	return doc
}
