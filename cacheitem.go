package simple_memory_cache

import (
	"sync"
	"time"
)

type CacheItem struct {
	sync.RWMutex
	//item's key
	key interface{}
	//item's data
	data interface{}
	//How long will the item live in the cache when not being accessed/kept alive.
	lifeSpan time.Duration
	//creation timestamp
	createOn time.Time
	//last access timestamp
	accessedOn time.Time
	//how often the item was accessed
	accessCount int64
	//callback method triggered right before removing the item from the cache
	aboutToExpire func(key interface{})
}

func NewCacheItem(key interface{}, lifeSpan time.Duration, data interface{}) *CacheItem {
	t := time.Now()
	return &CacheItem{
		key:           key,
		lifeSpan:      lifeSpan,
		createOn:      t,
		accessedOn:    t,
		accessCount:   0,
		aboutToExpire: nil,
		data:          data,
	}
}

func (item *CacheItem) KeepAlive() {
	item.Lock()
	defer item.Unlock()
	item.accessedOn = time.Now()
	item.accessCount++
}

func (item *CacheItem) AccessedOn() time.Time {
	item.RLock()
	defer item.RUnlock()
	return item.accessedOn
}

func (item *CacheItem) CreatedOn() time.Time {
	// immutable
	return item.createOn
}

func (item *CacheItem) AccessCount() int64 {
	item.RLock()
	defer item.RUnlock()
	return item.accessCount
}

func (item *CacheItem) Key() interface{} {
	// immutable
	return item.key
}

func (item *CacheItem) Data() interface{} {
	// immutable
	return item.data
}

func (item *CacheItem) SetAboutToExpireCallback(f func(interface{})) {
	item.Lock()
	defer item.Unlock()
	item.aboutToExpire = f
}
