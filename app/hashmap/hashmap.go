package hashmap

import (
	"sync"
	"time"
)

type HashMap interface {
	Set(key string, val string)
	Get(key string) (string, bool)
	Expire(key string, timeSeconds int) int
}

type Val struct {
	value       string
	setAt       time.Time
	expireAfter time.Duration
	expire      bool
}

type ConcurrentMap struct {
	mutex sync.RWMutex
	data  map[string]*Val
}

func Create() *ConcurrentMap {
	hashmap := ConcurrentMap{
		data: make(map[string]*Val),
	}
	return &hashmap
}
