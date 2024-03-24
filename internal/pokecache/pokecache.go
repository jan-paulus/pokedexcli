package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entries map[string]cacheEntry
	mux     *sync.Mutex
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       val,
	}
	c.mux.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	val := c.entries[key].val
	return val, val != nil
}

func (c *Cache) reapLoop(interval time.Duration) {
  ticker := time.NewTicker(interval)
	for range ticker.C {
    c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
  c.mux.Lock()
  defer c.mux.Unlock()
  for k, v := range c.entries {
    if v.createdAt.Before(now.Add(-last)) {
      delete(c.entries, k)
    }
  }
}


func NewCache(interval time.Duration) Cache {
	cache := Cache{
    entries: make(map[string]cacheEntry),
    mux: &sync.Mutex{},
	}

	go cache.reapLoop(interval)

	return cache
}
