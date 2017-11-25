package cache

import (
	"time"

	gocache "github.com/patrickmn/go-cache"
)

// Cache is the general app cache object
var Cache *gocache.Cache

// Boot initializes the cache
func Boot() {
	Cache = gocache.New(5*time.Minute, 10*time.Minute)
}

// Duration returns the default duration
func Duration() time.Duration {
	return gocache.DefaultExpiration
}

// Get an item from the cache. Returns the item or nil, and a bool indicating
func Get(key string) (interface{}, bool) {
	if Cache == nil {
		Boot()
	}
	return Cache.Get(key)
}

// Set adds an item to the cache, replacing any existing item.
func Set(key string, value interface{}, duration time.Duration) {
	if Cache == nil {
		Boot()
	}
	Cache.Set(key, value, duration)
}
