package fetcher

import (
	"log"
	"net/http"
	"os"

	"github.com/hit-bheda/web-crawler/internal/parser"
	"golang.org/x/net/html"
)

func FetchDocument(url string) {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 ...")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Faild to fetch page data :", err)
	}

	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal("Faild to fetch page data :", err)
	}

	file, err := os.Create("data.md")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	parser.TextParser(doc, file)
	parser.LinkParser(doc)
}
