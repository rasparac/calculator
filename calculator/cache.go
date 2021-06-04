package calculator

import (
	"time"

	"github.com/patrickmn/go-cache"
)

//Cache implements Cacher interface
type Cache struct {
	cache *cache.Cache
}

// NewCache returns a new Cache instance
func NewCache() *Cache {
	return &Cache{
		cache: cache.New(1*time.Minute, 1*time.Minute),
	}
}

// Get will return and value from cache by key if there is any
func (c *Cache) Get(key string) (float64, bool) {
	v, found := c.cache.Get(key)
	if !found {
		return 0, found
	}
	value, ok := v.(float64)
	if !ok {
		return 0, false
	}
	return value, found
}

// Set will set value in cache by key
func (c *Cache) Set(key string, v float64) {
	err := c.cache.Add(key, v, cache.DefaultExpiration)
	// ignore error for now, later check how to handle error
	if err != nil {
	}
}
