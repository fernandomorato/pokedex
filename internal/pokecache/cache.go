package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	rawData   []byte
}

type Cache struct {
	entries map[string]cacheEntry
	mutex   *sync.RWMutex
}

func NewCache(duration time.Duration) Cache {
	cache := Cache{
		entries: map[string]cacheEntry{},
		mutex:   &sync.RWMutex{},
	}
	go cache.reapLoop(duration)
	return cache
}

func (c *Cache) Add(key string, data []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		rawData:   data,
	}
}

func (c *Cache) Get(key string) (data []byte, ok bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	value, ok := c.entries[key]
	return value.rawData, ok
}

func (c *Cache) reapLoop(duration time.Duration) {
	ticker := time.NewTicker(duration)
	for range ticker.C {
		c.removeExpired(time.Now(), duration)
	}
}

func (c *Cache) removeExpired(now time.Time, duration time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	for k, v := range c.entries {
		if v.createdAt.Add(duration).Before(now) {
			delete(c.entries, k)
		}
	}
}
