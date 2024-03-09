package pokeCache

import (
	"sync"
	"time"
)

type Cache struct {
	cache    map[string]cacheEntry
	mux      *sync.RWMutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		cache:    map[string]cacheEntry{},
		mux:      &sync.RWMutex{},
		interval: interval,
	}

	go cache.ReapLoop()
	return cache
}

func (ch *Cache) Add(key string, val []byte) {
	ch.mux.Lock()
	defer ch.mux.Unlock()

	ch.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (ch *Cache) Get(key string) ([]byte, bool) {
	ch.mux.RLock()
	defer ch.mux.RUnlock()

	entry, ok := ch.cache[key]

	if !ok {
		return []byte{}, false
	}

	return entry.val, true

}

func (ch *Cache) ReapLoop() {
	ticker := time.NewTicker(ch.interval)
	// not sure this will ever run
	defer ticker.Stop()
	for {
		// wait for a tick from the channel
		<-ticker.C

		// lock as we go through it so R and W are blocked since we might delete?
		ch.mux.Lock()
		for key, entry := range ch.cache {
			if time.Since(entry.createdAt) > ch.interval {
				delete(ch.cache, key)
			}
		}
		ch.mux.Unlock()
	}
}
