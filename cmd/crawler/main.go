package main

import (
	"context"
	"flag"
	"log"
	"time"

	"Concurent-WebCrawler/internal/crawler"
	"Concurent-WebCrawler/internal/fetcher"
	"Concurent-WebCrawler/internal/orchestrator"
)

func main() {

	startURL := flag.String(
		"url",
		"",
		"starting URL",
	)

	timeout := flag.Duration(
		"timeout",
		10*time.Second,
		"http timeout",
	)

	depth := flag.Int(
		"depth",
		1,
		"maximum crawl depth",
	)

	workers := flag.Int(
		"workers",
		5,
		"number of concurrent workers",
	)

	flag.Parse()

	if *startURL == "" {
		log.Fatal("url flag is required")
	}

	ctx := context.Background()

	f := fetcher.New(*timeout)

	c := crawler.New(
		f,
		*depth,
	)

	pool := orchestrator.NewWorkerPool(
		*workers,
	)

	pool.Start(
		ctx,
		func(
			ctx context.Context,
			job crawler.Job,
		) {

			c.CrawlJob(
				ctx,
				job,
				pool.Jobs,
			)
		},
	)

	c.AddJob()

	pool.Jobs <- crawler.Job{
		URL:   *startURL,
		Depth: 0,
	}

	c.Wait()

	close(pool.Jobs)

	pool.Wait()
}
