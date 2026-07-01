package t1

import "sync"

type BondedQueue struct {
	data    []any
	maxSize int
	lock    sync.Mutex
	cond    sync.Cond
}

func NewBondedQueue(maxSize int) *BondedQueue {
	queue := BondedQueue{
		data:    make([]any, 0, maxSize),
		maxSize: maxSize,
		lock:    sync.Mutex{},
		cond:    *sync.NewCond(&sync.Mutex{}),
	}

	return &queue
}

func (b *BondedQueue) Put(value any) {
	if len(b.data) == b.maxSize {
		b.cond.Wait()
	}

	b.lock.Lock()
	defer b.lock.Unlock()

	b.data = append(b.data, value)
	if len(b.data) == 1 {
		// не совсем понятно что тут лучше использовать broadcast, или signal...
		// по идее broadcast ибо несколько горутин могут хотеть записать значение,
		// но с другой стороны, как только новое значение запишется cond снова заблокирует исполнение...
		b.cond.Signal()
	}
}

func (b *BondedQueue) Get() any {
	if len(b.data) == 0 {
		b.cond.Wait()
	}

	b.lock.Lock()
	defer b.lock.Unlock()

	value := b.data[0]
	b.data = b.data[1:]

	if len(b.data) == b.maxSize-1 {
		b.cond.Signal()
	}

	return value
}

func (b *BondedQueue) Length() int {
	return len(b.data)
}

func (b *BondedQueue) Shutdown() {

}
