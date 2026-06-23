package cache

import (
	"encoding/json"
	"errors"
	"time"
)

type cacheUnit struct {
	value any
	timer *time.Timer
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
	timer := time.NewTimer(ttl)
	c.data[key] = cacheUnit{
		value: value,
		timer: timer,
	}
	go func() {
		<-timer.C
		delete(c.data, key)
	}()
}

func (c *Cache) Clear() {
	for _, val := range c.data {
		val.timer.Stop()
	}
}

func (c *Cache) Get(key string) (any, bool) {
	data, exists := c.data[key]
	return data.value, exists
}

func (c *Cache) Delete(key string) {
	data, exists := c.data[key]
	if exists {
		data.timer.Stop()
	}
}

func (c *Cache) Exists(key string) bool {
	_, exists := c.data[key]
	return exists
}

func (c *Cache) ToJSON(key string) ([]byte, error) {
	data := make(map[string]any, len(c.data))

	for key, value := range c.data {
		data[key] = value.value
	}

	jsonString, err := json.Marshal(data)

	return jsonString, err
}

func GetAs[T any](c *Cache, key string) (T, error) {
	var zero T
	return zero, errors.New("Not implemented")
}
