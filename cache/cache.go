package cache

import (
	"time"

	cache "github.com/patrickmn/go-cache"
)

// Cache is the general app cache object
var Cache *cache.Cache

// Boot initializes the cache
func Boot() {
	Cache = cache.New(5*time.Minute, 10*time.Minute)
}

// Duration returns the default duration
func Duration() time.Duration {
	return cache.DefaultExpiration
}
