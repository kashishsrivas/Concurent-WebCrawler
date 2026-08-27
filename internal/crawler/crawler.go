package crawler

import (
	"context"
	"fmt"
	"sync"

	"Concurent-WebCrawler/internal/fetcher"
	"Concurent-WebCrawler/internal/parser"
)

type Crawler struct {
	fetcher  *fetcher.Fetcher
	visited  *Visited
	maxDepth int
	wg       sync.WaitGroup

	maxPages  int
	pageCount int
}

func New(
	f *fetcher.Fetcher,
	maxDepth int,
) *Crawler {

	return &Crawler{
		fetcher:  f,
		visited:  NewVisited(),
		maxDepth: maxDepth,
		maxPages: 1000,
	}
}

func (c *Crawler) Crawl(
	ctx context.Context,
	url string,
	depth int,
) {

	if depth > c.maxDepth {
		return
	}

	if c.visited.Seen(url) {
		return
	}

	c.visited.Add(url)

	fmt.Printf(
		"\n[Depth %d] Crawling: %s\n",
		depth,
		url,
	)

	body, _, err := c.fetcher.Fetch(
		ctx,
		url,
	)

	if err != nil {
		fmt.Println("Fetch Error:", err)
		return
	}

	links, err := parser.ExtractLinks(
		body,
		url,
	)

	if err != nil {
		fmt.Println("Parse Error:", err)
		return
	}

	fmt.Printf(
		"[Depth %d] Found %d links\n",
		depth,
		len(links),
	)

	for _, link := range links {
		c.Crawl(
			ctx,
			link,
			depth+1,
		)
	}
}

// Worker Pool Version
func (c *Crawler) CrawlJob(
	ctx context.Context,
	job Job,
	jobs chan<- Job,
) {

	defer c.DoneJob()

	if job.Depth > c.maxDepth {
		return
	}

	if c.visited.Seen(job.URL) {
		return
	}

	c.visited.Add(job.URL)

	fmt.Printf(
		"\n[Worker] Crawling: %s (Depth %d)\n",
		job.URL,
		job.Depth,
	)

	body, _, err := c.fetcher.Fetch(
		ctx,
		job.URL,
	)

	if err != nil {
		fmt.Println("Fetch Error:", err)
		return
	}

	links, err := parser.ExtractLinks(
		body,
		job.URL,
	)

	if err != nil {
		fmt.Println("Parse Error:", err)
		return
	}

	fmt.Printf(
		"[Depth %d] Found %d links\n",
		job.Depth,
		len(links),
	)
	for _, link := range links {

		if job.Depth+1 > c.maxDepth {
			continue
		}

		if !c.visited.Seen(link) {

			c.AddJob()

			jobs <- Job{
				URL:   link,
				Depth: job.Depth + 1,
			}
		}
	}

}

func (c *Crawler) AddJob() {
	c.wg.Add(1)
}

func (c *Crawler) DoneJob() {
	c.wg.Done()
}

func (c *Crawler) Wait() {
	c.wg.Wait()
}
