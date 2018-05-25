package cache

import (
	"log"
	"sync"
	"time"

	gocache "github.com/patrickmn/go-cache"
)

// Cache is the general app cache object
var cache *gocache.Cache
var cacheInitializer sync.Once

// Boot initializes the cache
func Boot() {
	cacheInitializer.Do(func() {
		log.Println("Booting cache")
		cache = gocache.New(5*time.Minute, 10*time.Minute)
	})
}

// Duration returns the default duration
func Duration() time.Duration {
	return gocache.DefaultExpiration
}

// Get an item from the cache. Returns the item or nil, and a bool indicating
func Get(key string) (interface{}, bool) {
	if cache == nil {
		Boot()
	}
	return cache.Get(key)
}

// Set adds an item to the cache, replacing any existing item.
func Set(key string, value interface{}, duration time.Duration) {
	if cache == nil {
		Boot()
	}
	cache.Set(key, value, duration)
}

func Clear(key string) {
	if cache == nil {
		Boot()
	}
	cache.Delete(key)
}
