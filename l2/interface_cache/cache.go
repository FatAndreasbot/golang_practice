package cache

import (
	"encoding/json"
	"errors"
	"time"
)

// TODO no goroutines
type cacheUnit struct {
	value any
	timer time.Time
}

type Cache struct {
	data map[string]cacheUnit
}

func NewCache() Cache {
	return Cache{
		data: make(map[string]cacheUnit),
	}
}

func (c *Cache) Set(key string, value any, ttl time.Duration) {
	timer := time.Now().Add(-ttl)

	c.data[key] = cacheUnit{
		value: value,
		timer: timer,
	}
}

func (c *Cache) Clear() {
	for key := range c.data {
		delete(c.data, key)
	}
}

func (c *Cache) Get(key string) (any, bool) {
	data, exists := c.data[key]
	if !exists {
		return data.value, exists
	}

	if time.Now().Before(data.timer) {
		delete(c.data, key)
		return struct{}{}, false
	}

	return data.value, true
}

func (c *Cache) Delete(key string) {
	_, exists := c.data[key]
	if exists {
		delete(c.data, key)
	}
}

func (c *Cache) Exists(key string) bool {
	data, exists := c.data[key]

	if !exists {
		return false
	}

	if time.Now().Before(data.timer) {
		delete(c.data, key)
		return false
	}

	return exists
}

func (c *Cache) ToJSON() ([]byte, error) {
	data := make(map[string]any)

	for key, value := range c.data {
		if !c.Exists(key) {
			continue
		}
		data[key] = value.value
	}

	jsonString, err := json.Marshal(data)

	return jsonString, err
}

func GetAs[T any](c *Cache, key string) (val T, err error) {
	var zero T

	data, exists := c.data[key]

	if !exists {
		return zero, errors.New("empty value")
	}
	val, ok := data.value.(T)
	if ok {
		return val, nil
	}
	return zero, errors.New("casting error")
}
