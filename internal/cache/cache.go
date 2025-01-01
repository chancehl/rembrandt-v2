package cache

import (
	"sync"
	"time"
)

type InMemoryCacheItem struct {
	value  interface{}
	expiry time.Time
}

type InMemoryCache struct {
	data  map[string]InMemoryCacheItem
	mutex sync.RWMutex
}

// Creates a new InMemoryCache
func NewInMemoryCache() *InMemoryCache {
	return &InMemoryCache{
		data: make(map[string]InMemoryCacheItem),
	}
}

// Sets a cached value
func (c *InMemoryCache) Set(key string, value interface{}, ttl time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.data[key] = InMemoryCacheItem{
		value:  value,
		expiry: time.Now().Add(ttl),
	}
}

// Gets a cached value
//
// Note: if the user requests an item that is past the expiry it will be deleted from the cache
func (c *InMemoryCache) Get(key string) (interface{}, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	item, ok := c.data[key]
	if !ok {
		return nil, false
	}

	if item.expiry.Before(time.Now()) {
		delete(c.data, key)
		return nil, false
	}
	return item.value, true
}

// Deletes a cached value
func (c *InMemoryCache) Delete(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	delete(c.data, key)
}

// Clears all cached values
func (c *InMemoryCache) Clear() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.data = make(map[string]InMemoryCacheItem)
}
