package fetcher

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Fetcher struct {
	client *http.Client
}

func New(timeout time.Duration) *Fetcher {
	return &Fetcher{
		client: &http.Client{
			Timeout: timeout,
		},
	}
}

func (f *Fetcher) Fetch(
	ctx context.Context,
	targetURL string,
) ([]byte, int, error) {

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		targetURL,
		nil,
	)

	if err != nil {
		return nil, 0, err
	}

	req.Header.Set(
		"User-Agent",
		"GoCrawler/1.0",
	)

	resp, err := f.client.Do(req)

	if err != nil {
		return nil, 0, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, 0, err
	}

	if resp.StatusCode >= 400 {
		return nil,
			resp.StatusCode,
			fmt.Errorf("bad status code: %d", resp.StatusCode)
	}

	return body, resp.StatusCode, nil
}
