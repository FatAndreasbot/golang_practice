package obj_cache

import (
	"bytes"
	"encoding/json"
	"sync"
	"time"
)

type cacheUnit struct {
	data  any
	timer *time.Timer
}

type Cache struct {
	data map[string]cacheUnit
	mu   sync.RWMutex
	pool sync.Pool
	ttl  time.Duration
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	unit, ok := c.data[key]
	if ok {
		unit.timer.Stop()
		delete(c.data, key)
	}

}

func (c *Cache) Set(key string, value any) {
	newCacheUnit := cacheUnit{
		data: value,
		timer: time.AfterFunc(c.ttl, func() {
			c.mu.Lock()
			defer c.mu.Unlock()
			delete(c.data, key)
		}),
	}
	c.data[key] = newCacheUnit
}

func (c *Cache) Get(key string) (any, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	data, ok := c.data[key]
	if ok {
		return data.data, ok
	}
	return nil, ok
}

func (c *Cache) ToJSON() ([]byte, error) {
	totalBuffer := c.pool.Get().(*bytes.Buffer)
	fieldBuffer := c.pool.Get().(*bytes.Buffer)
	defer c.pool.Put(totalBuffer)
	defer c.pool.Put(fieldBuffer)
	defer totalBuffer.Reset()
	defer fieldBuffer.Reset()

	err := totalBuffer.WriteByte('{')

	for key, value := range c.data {
		enc := json.NewEncoder(fieldBuffer)
		err = enc.Encode(value.data)
		if err != nil {
			return nil, err
		}
		fieldBuffer.Truncate(fieldBuffer.Len() - 1)

		totalBuffer.WriteByte('"')
		totalBuffer.Write([]byte(key))
		totalBuffer.Write([]byte("\":"))
		_, err = totalBuffer.Write(fieldBuffer.Bytes())
		if err != nil {
			return nil, err
		}
		totalBuffer.WriteByte(',')

		fieldBuffer.Reset()
	}
	totalBuffer.Truncate(totalBuffer.Len() - 1)
	totalBuffer.WriteByte('}')

	return totalBuffer.Bytes(), nil
}

func NewObjectCache(ttl time.Duration) Cache {
	return Cache{
		data: make(map[string]cacheUnit),
		mu:   sync.RWMutex{},
		ttl:  ttl,
		pool: sync.Pool{
			New: func() any { return bytes.NewBuffer([]byte{}) },
		},
	}
}
