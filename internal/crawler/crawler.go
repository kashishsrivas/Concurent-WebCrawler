package crawler

import (
	"context"
	"fmt"

	"Concurent-WebCrawler/internal/fetcher"
	"Concurent-WebCrawler/internal/parser"
)

type Crawler struct {
	fetcher  *fetcher.Fetcher
	visited  *Visited
	maxDepth int
}

func New(
	f *fetcher.Fetcher,
	maxDepth int,
) *Crawler {

	return &Crawler{
		fetcher:  f,
		visited:  NewVisited(),
		maxDepth: maxDepth,
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
