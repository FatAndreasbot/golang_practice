package t1_test

import (
	"math/rand"
	"t1"
	"testing"
	"time"
)

func TestQueue(t *testing.T) {
	queue := t1.NewBondedQueue(3)

	go func() {
		for range 5 {
			queue.Put(rand.Int())
		}
	}()

	time.Sleep(time.Millisecond * 100)
	if queue.Length() > 3 {
		t.Error("the limit is not working(((")
	}

	go func() {
		for range 3 {
			_ = queue.Get()
		}
	}()

	time.Sleep(time.Millisecond * 100)
	if queue.Length() != 2 {
		t.Error("the queue is not filling")
	}

}
