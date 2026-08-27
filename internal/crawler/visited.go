package crawler

import "sync"

type Visited struct {
	mu   sync.RWMutex
	urls map[string]bool
}

func NewVisited() *Visited {
	return &Visited{
		urls: make(map[string]bool),
	}
}

func (v *Visited) Seen(url string) bool {
	v.mu.RLock()
	defer v.mu.RUnlock()

	return v.urls[url]
}

func (v *Visited) Add(url string) {
	v.mu.Lock()
	defer v.mu.Unlock()

	v.urls[url] = true
}
