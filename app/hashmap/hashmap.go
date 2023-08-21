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
func (c *ConcurrentMap) Set(key string, value string) {
	c.mutex.Lock()
	c.data[key] = &Val{
		value:       value,
		setAt:       time.Now(),
		expireAfter: 0,
		expire:      false,
	}
	c.mutex.Unlock()
}
func (c *ConcurrentMap) Get(key string) (string, bool) {
	c.mutex.RLock()
	valueItem, exists := c.data[key]
	if !exists {
		c.mutex.RUnlock()
		return "", false
	}

	if valueItem.expire && time.Now().Sub(valueItem.setAt) > valueItem.expireAfter {
		c.mutex.RUnlock()
		return "", false
	}
	c.mutex.RUnlock()
	return valueItem.value, exists
}
