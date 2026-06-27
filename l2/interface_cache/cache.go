package cache

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

// TODO no goroutines
type cacheUnit struct {
	value any
	timer *time.Timer
}

type Cache struct {
	data map[string]cacheUnit
	lock sync.Mutex
}

func NewCache() Cache {
	return Cache{
		data: make(map[string]cacheUnit),
		lock: sync.Mutex{},
	}
}

func (c *Cache) Set(key string, value any, ttl time.Duration) {
	timer := time.AfterFunc(ttl, func() {
		c.lock.Lock()
		defer c.lock.Unlock()
		fmt.Printf("deleting %s\n", key)
		delete(c.data, key)
	})

	c.data[key] = cacheUnit{
		value: value,
		timer: timer,
	}
}

func (c *Cache) Clear() {
	c.lock.Lock()
	defer c.lock.Unlock()

	for key, val := range c.data {
		val.timer.Stop()
		delete(c.data, key)
	}
}

func (c *Cache) Get(key string) (any, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()

	data, exists := c.data[key]
	return data.value, exists
}

func (c *Cache) Delete(key string) {
	c.lock.Lock()
	defer c.lock.Unlock()

	data, exists := c.data[key]
	if exists {
		data.timer.Stop()
		delete(c.data, key)
	}
}

func (c *Cache) Exists(key string) bool {
	_, exists := c.data[key]
	return exists
}

func (c *Cache) ToJSON() ([]byte, error) {
	c.lock.Lock()
	defer c.lock.Unlock()

	data := make(map[string]any, len(c.data))

	for key, value := range c.data {
		data[key] = value.value
	}

	jsonString, err := json.Marshal(data)

	return jsonString, err
}

func GetAs[T any](c *Cache, key string) (val T, err error) {
	c.lock.Lock()
	defer c.lock.Unlock()

	defer func() {
		r := recover()
		if r != nil {
			err = fmt.Errorf("casting panic %#v", r)
		}
	}()

	data, exists := c.data[key]
	if exists {
		val = data.value.(T)
	}
	return val, err
}
