package cache

import (
	"github.com/prometheus/common/log"
	"sync"
	"time"
)

type Cache struct {
	ttl             time.Duration
	cleanupInterval time.Duration
	items           sync.Map
}

func NewCache(ttl time.Duration, cleanupInterval time.Duration) *Cache {
	cache := &Cache{
		ttl:             ttl,
		cleanupInterval: cleanupInterval,
	}
	cache.startCleanupTimer()
	return cache
}

func (c *Cache) Store(key interface{}, data string) {
	item := Item{data: data, expires: time.Now().Add(c.ttl)}
	c.items.Store(key, item)
}

func (c *Cache) Load(key interface{}) (value Item, ok bool) {
	val, ok := c.items.Load(key)
	if ok {
		return val.(Item), ok
	}
	return Item{}, false
}

func (c *Cache) HasElement(key string) bool {
	_, ok := c.items.Load(key)
	return ok
}

func (c *Cache) startCleanupTimer() {

	duration := c.cleanupInterval
	if duration < time.Second {
		duration = time.Second
	}
	ticker := time.Tick(duration)
	go (func() {
		for {
			select {
			case <-ticker:
				c.cleanup()
			}
		}
	})()
}

func (c *Cache) cleanup() {
	log.Info("Started cleanup")
	c.items.Range(func(key, value interface{}) bool {
		item := value.(Item)
		if item.expired() {
			log.Info("delete", key)
			c.items.Delete(key)
		}
		return true
	})
}

func (c *Cache) Count() int {
	length := 0

	c.items.Range(func(_, _ interface{}) bool {
		length++
		return true
	})

	return length
}
