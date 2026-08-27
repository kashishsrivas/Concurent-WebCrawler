package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"Concurent-WebCrawler/internal/crawler"
	"Concurent-WebCrawler/internal/fetcher"
	"Concurent-WebCrawler/internal/orchestrator"
	"Concurent-WebCrawler/internal/storage"
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

	ctx, cancel := context.WithCancel(
		context.Background(),
	)
	defer cancel()
	signalChan := make(chan os.Signal, 1)

	signal.Notify(
		signalChan,
		os.Interrupt,
		syscall.SIGTERM,
	)

	go func() {
		<-signalChan

		log.Println("Shutdown signal received...")
		cancel()
	}()

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

	if err := storage.SaveJSON(
		"output/crawl.json",
		c.Results(),
	); err != nil {
		log.Println(err)
	}

	if err := storage.SaveCSV(
		"output/crawl.csv",
		c.Results(),
	); err != nil {
		log.Println(err)
	}

	close(pool.Jobs)

	log.Println("================================")
	log.Printf("Pages Crawled : %d\n", c.PagesCrawled())
	log.Printf("Workers Used  : %d\n", *workers)
	log.Printf("Max Depth     : %d\n", *depth)
	log.Println("================================")

	pool.Wait()

	log.Println("Crawler stopped successfully.")

}
