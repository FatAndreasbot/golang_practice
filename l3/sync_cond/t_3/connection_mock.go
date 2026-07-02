package t3

import (
	"errors"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

var NotFoundError error = errors.New("value was not found")
var mockData map[string]any

var doOnce sync.Once

type Connection struct {
	ID   int
	data *map[string]any
	lock sync.RWMutex
}

func (c *Connection) GetData(key string) (any, error) {
	c.lock.RLock()
	defer c.lock.Unlock()

	data, ok := (*c.data)[key]
	if !ok {
		return nil, NotFoundError
	}
	time.Sleep(100 * time.Millisecond)
	return data, nil
}

func (c *Connection) SetData(key string, value any) {
	c.lock.RLock()
	defer c.lock.Unlock()

	time.Sleep(100 * time.Millisecond)
	(*c.data)[key] = value
}

func createAMockConnection(connID int) *Connection {
	connection := Connection{
		ID:   connID,
		data: &mockData,
		lock: sync.RWMutex{},
	}

	doOnce.Do(func() {
		mockData = make(map[string]any)
		for i := range 5 {
			mockData[strconv.FormatInt(int64(i), 10)] = rand.Int()
		}
	})

	return &connection
}
