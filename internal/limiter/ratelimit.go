package limiter

import "time"

type RateLimiter struct {
	ticker *time.Ticker
}

func NewRateLimiter(
	interval time.Duration,
) *RateLimiter {

	return &RateLimiter{
		ticker: time.NewTicker(interval),
	}
}

func (r *RateLimiter) Wait() {
	<-r.ticker.C
}
