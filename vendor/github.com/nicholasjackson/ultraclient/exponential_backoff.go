package ultraclient

import (
	"time"

	"github.com/eapache/go-resiliency/retrier"
)

type ExponentialBackoff struct {
	cache []time.Duration
}

func (e *ExponentialBackoff) Create(retries int, delay time.Duration) []time.Duration {
	if e.cache == nil {
		e.cache = retrier.ExponentialBackoff(retries, delay)
	}

	return e.cache
}
