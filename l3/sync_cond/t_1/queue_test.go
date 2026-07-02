package t1_test

import (
	"math/rand"
	"t1"
	"testing"
	"time"
)

func TestQueue(t *testing.T) {
	queue := t1.NewBondedQueue[int](3)

	go func() {
		for i := range 5 {
			go queue.Put(i)
		}
	}()

	time.Sleep(time.Millisecond * 300)
	if queue.Length() != 3 {
		t.Error("the queue does not fill up(((")
	}

	if queue.Length() > 3 {
		t.Error("the limit is not working(((")
	}
	t.Log("tested insertion")

	go func(t *testing.T) {
		for range 3 {
			val := queue.Get()
			t.Logf("value: %d", val)
		}
	}(t)

	time.Sleep(time.Millisecond * 300)
	t.Log(queue.Length())
	if queue.Length() != 2 {
		t.Error("the queue is not filling")
	}

	queue.Shutdown()
	go func() {
		queue.Put(rand.Int())
	}()
	time.Sleep(time.Millisecond * 100)
	t.Log(queue.Length())
	if queue.Length() > 2 {
		t.Error("the queue is fillign after shutdown")
	}

}
