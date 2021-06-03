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
	value, ok := v.(float64)
	if !ok {
		return 0, false
	}
	return value, found
}

func (c *Cache) Set(key string, v float64) {
	err := c.cache.Add(key, v, cache.DefaultExpiration)
	if err != nil {
	}
}
