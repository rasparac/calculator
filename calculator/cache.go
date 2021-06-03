package calculator

import (
	"time"

	"github.com/patrickmn/go-cache"
)

type Cache struct {
	cache *cache.Cache
}

func NewCache() *Cache {
	return &Cache{
		cache: cache.New(1*time.Minute, 1*time.Minute),
	}
}

func (c *Cache) Get(key string) (float64, bool) {
	v, found := c.cache.Get(key)
	if !found {
		return 0, found
	}
	return v.(float64), found
}

func (c *Cache) Set(key string, v float64) {
	c.cache.Add(key, v, cache.DefaultExpiration)
}
