package storage

type CrawlResult struct {
	URL   string `json:"url"`
	Depth int    `json:"depth"`
}
