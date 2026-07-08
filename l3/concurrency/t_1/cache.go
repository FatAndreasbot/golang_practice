package t1

import "sync"

type Cache[K comparable, T any] struct {
	data map[K]T
	lock sync.RWMutex
}

func NewCache[K comparable, T any]() *Cache[K, T] {
	return &Cache[K, T]{
		data: make(map[K]T),
		lock: sync.RWMutex{},
	}
}

func (c *Cache[K, T]) Set(key K, value T) {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.data[key] = value
}

func (c *Cache[K, T]) Get(key K) (*T, bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	data, ok := c.data[key]

	return &data, ok
}

func (c *Cache[K, T]) Delete(key K) {
	c.lock.Lock()
	defer c.lock.Unlock()

	delete(c.data, key)
}
