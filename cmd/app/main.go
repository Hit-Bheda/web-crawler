package main

import (
	"github.com/hit-bheda/web-crawler/internal/fetcher"
	"github.com/hit-bheda/web-crawler/internal/logger"
	"github.com/hit-bheda/web-crawler/internal/parser"
	"github.com/hit-bheda/web-crawler/internal/writer"
)

func main() {
	log := logger.New()
	log.Info().Msg("Starting the executiuon!")
	url := "https://stackoverflow.com/questions/49130026/writing-json-objects-to-file"
	doc := fetcher.FetchDocument(url, log)

	docTitle, err := parser.GetTitle(doc)
	if err != nil {
		log.Error().Err(err).Str("url", url).Msg("Failed to get title ")
	}

	textContent := parser.TextParser(doc)

	writer.WriteDoc(url, docTitle, textContent, log)

	parser.LinkParser(doc, url, log)
}
