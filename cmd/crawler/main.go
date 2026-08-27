package main

import (
	"context"
	"flag"
	"log"
	"time"

	"Concurent-WebCrawler/internal/crawler"
	"Concurent-WebCrawler/internal/fetcher"
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

	c.Crawl(
		ctx,
		*startURL,
		0,
	)
}
