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
	cacheMap map[string]cacheEntry
	mux      *sync.Mutex
	Interval time.Duration
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{}
	cache.cacheMap = map[string]cacheEntry{}
	cache.mux = &sync.Mutex{}
	cache.Interval = interval
	go cache.reapLoop()
	return cache
}

func (cache Cache) Add(key string, val []byte) {
	cache.mux.Lock()
	entry := cacheEntry{}
	entry.createdAt = time.Now()
	entry.val = val
	cache.cacheMap[key] = entry
	cache.mux.Unlock()
}

func (cache Cache) Get(key string) ([]byte, bool) {
	cache.mux.Lock()
	entry, ok := cache.cacheMap[key]
	val := entry.val
	cache.mux.Unlock()
	return val, ok
}

func (cache Cache) reapLoop() {
	interval := cache.Interval
	for {
		time.Sleep(interval)
		cache.mux.Lock()
		for key := range cache.cacheMap {
			if time.Since(cache.cacheMap[key].createdAt) > interval {
				delete(cache.cacheMap, key)
			}
		}
		cache.mux.Unlock()
	}
}
