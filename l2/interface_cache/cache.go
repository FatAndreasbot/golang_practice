package cache

import (
	"encoding/json"
	"fmt"
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
	timer := time.AfterFunc(ttl, func() {
		fmt.Printf("deleting %s\n", key)
		delete(c.data, key)
	})

	c.data[key] = cacheUnit{
		value: value,
		timer: timer,
	}
}

func (c *Cache) Clear() {
	for key, val := range c.data {
		val.timer.Stop()
		delete(c.data, key)
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
		delete(c.data, key)
	}
}

func (c *Cache) Exists(key string) bool {
	_, exists := c.data[key]
	return exists
}

func (c *Cache) ToJSON() ([]byte, error) {
	data := make(map[string]any, len(c.data))

	for key, value := range c.data {
		data[key] = value.value
	}

	jsonString, err := json.Marshal(data)

	return jsonString, err
}

func GetAs[T any](c *Cache, key string) (T, error) {
	var err error
	var val T
	defer func() {
		r := recover()
		err = fmt.Errorf("casting panic %#v", r)
	}()

	val = c.data[key].value.(T)
	return val, err
}
