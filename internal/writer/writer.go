package writer

import (
	"encoding/json"
	"os"

	"github.com/hit-bheda/web-crawler/internal/hash"
	"github.com/hit-bheda/web-crawler/internal/types"
	"github.com/rs/zerolog"
)

func WriteDoc(url string, docTitle string, textContent string, log zerolog.Logger) {
	fileName := hash.HashFilename(url)
	file, err := os.Create(string("docs/"+string(fileName)) + ".json")
	if err != nil {
		log.Error().Err(err).Str("url", url).Msg("Failed write data in file!")
	}
	defer file.Close()
	document := []types.Document{{
		Id:      string(fileName),
		URL:     url,
		Title:   docTitle,
		Content: string(textContent),
	}}
	jsonData, _ := json.MarshalIndent(document, "", "\t")
	file.Write(jsonData)
}
