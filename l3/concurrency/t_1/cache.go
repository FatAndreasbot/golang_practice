package t1

import "sync"

type Cache struct{
	data map[string]string
	lock sync.RWMutex
}

func NewCache() Cache{
	return Cache{
		data: make(map[string]string),
		lock: sync.RWMutex{},
	}
}

func (c *Cache) Set (key string, value string){
	c.lock.Lock()
	c.data[key] = value
	c.lock.Unlock()
}

func (c *Cache) Get (key string) (string, bool){
	c.lock.RLock()
	data, ok := c.data[key]
	c.lock.RUnlock()
	return data, ok
}
