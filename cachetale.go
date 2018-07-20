package simple_memory_cache

import (
	"sync"
	"time"
)

type CacheTable struct {
	sync.RWMutex
	//The table's name.
	name         string
	items        map[interface{}]*CacheItem
	cleanupTimer *time.Timer

}
