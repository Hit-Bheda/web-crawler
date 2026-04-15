package main

import (
	"context"
	"os"

	"github.com/hit-bheda/web-crawler/internal/fetcher"
	"github.com/hit-bheda/web-crawler/internal/logger"
	"github.com/hit-bheda/web-crawler/internal/parser"
	"github.com/hit-bheda/web-crawler/internal/postgresql"
	"github.com/hit-bheda/web-crawler/internal/queue"
	"github.com/hit-bheda/web-crawler/internal/redis"
	"github.com/hit-bheda/web-crawler/internal/writer"
)

func main() {
	log := logger.New()
	log.Info().Msg("Starting the executiuon!")
	ctx := context.Background()
	rdb := redis.InitRedis(ctx)
	conn := postgresql.ConnectDB(log, ctx)
	defer conn.Close(ctx)

	for {
		url, err := queue.Dequeue(ctx, rdb)
		if err != nil {
			log.Error().Err(err).Str("url", url).Msg("Queue is empty!")
			os.Exit(1)
		}

		err = postgresql.InsertUrl(ctx, conn, url)
		if err != nil {
			log.Error().Err(err).Str("url", url).Msg("Url already exists in db!")
			continue
		}

		log.Info().Str("url", url).Msg("Crawling the site")
		doc, err := fetcher.FetchDocument(url, log)
		if err != nil {
			log.Error().Err(err).Str("url", url).Msg("Failed to fetch document")
			continue
		}

		docTitle, err := parser.GetTitle(doc)
		if err != nil {
			log.Error().Err(err).Str("url", url).Msg("Failed to get title ")
		}

		textContent := parser.TextParser(doc)

		writer.WriteDoc(url, docTitle, textContent, log)

		urls, _ := parser.LinkParser(doc, url, log)

		log.Info().Str("url", url).Msg("Sucessfully crawled the page!")
		for _, link := range urls {
			exists, _ := rdb.SIsMember(ctx, "visited", link).Result()
			if exists == false {
				rdb.SAdd(ctx, "visited", link)
				queue.Enqueue(ctx, rdb, link)
			}

		}
	}
}
