package main

import (
	"fmt"
	"sync"
	"time"
)

func getPurchasesByID(id int, url string) string {
	time.Sleep(3 * time.Second)

	return "response"
}

func main() {
	url := "https://www.wildberries.ru/getPurchasesByID?id=%d"
	responses := make([]string, 100_000, 100_000)
	jobs := make(chan int, 100_000)
	var working sync.WaitGroup

	for range 300 {
		go func() {
			for id := range jobs {
				responses[id] = getPurchasesByID(id, url)
				working.Done()
			}
		}()
	}

	for jobID := range 100_000 {
		jobs <- jobID
		working.Add(1)
	}

	working.Wait()
	close(jobs)

	for _, response := range responses {
		fmt.Println(response)
	}
}

// -------------------------------------------------------

type Chat struct {
	messages []string
	lock     sync.RWMutex
}

func NewChat() Chat {
	return Chat{
		messages: make([]string, 0),
		lock:     sync.RWMutex{},
	}
}

func (c *Chat) Receive(id int) (string, error) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	exists := id < len(c.messages)
	if !exists {
		return "", fmt.Errorf("message with id %d does not exists", id)
	}
	msg := c.messages[id]
	return msg, nil
}

func (c *Chat) Send(msg string) int {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.messages = append(c.messages, msg)
	return len(c.messages) - 1
}
