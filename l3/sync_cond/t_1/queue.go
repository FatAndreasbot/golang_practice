package t1

import "sync"

type BondedQueue[T any] struct {
	data    []T
	maxSize int
	cond    sync.Cond
}

func NewBondedQueue[T any](maxSize int) *BondedQueue[T] {
	queue := BondedQueue[T]{
		data:    make([]T, 0, maxSize),
		maxSize: maxSize,
		cond:    *sync.NewCond(&sync.Mutex{}),
	}

	return &queue
}

func (b *BondedQueue[T]) Put(value T) {
	b.cond.L.Lock()
	defer b.cond.L.Unlock()

	if len(b.data) >= b.maxSize {
		b.cond.Wait()
	}

	b.data = append(b.data, value)
	if len(b.data) == 1 {
		// не совсем понятно что тут лучше использовать broadcast, или signal...
		// по идее broadcast ибо несколько горутин могут хотеть записать значение,
		// но с другой стороны, как только новое значение запишется cond снова заблокирует исполнение...
		b.cond.Signal()
	}
}

func (b *BondedQueue[T]) Get() T {
	b.cond.L.Lock()
	defer b.cond.L.Unlock()

	if len(b.data) == 0 {
		b.cond.Wait()
	}

	value := b.data[0]
	b.data = b.data[1:]

	if len(b.data) == b.maxSize-1 {
		b.cond.Signal()
	}

	return value
}

func (b *BondedQueue[T]) Length() int {
	return len(b.data)
}

func (b *BondedQueue[T]) Shutdown() {
	b.maxSize = 0
	// простой и рабоий способ запретить ввод новых данных в очередь
}

func (b *BondedQueue[T]) IsClosed() bool {
	return b.maxSize == 0
}
