package t2

import (
	"errors"
	"sync"
)

type Restaurant struct {
	cond          *sync.Cond
	tableCount    int
	occupiedCount int
}

func NewReastaurant(tableCount int) *Restaurant {
	return &Restaurant{
		tableCount:    tableCount,
		occupiedCount: 0,
		cond:          sync.NewCond(&sync.Mutex{}),
	}
}

func (r *Restaurant) OccupyTable() {
	r.cond.L.Lock()
	defer r.cond.L.Unlock()

	if r.occupiedCount >= r.tableCount {
		r.cond.Wait()
	}

	r.occupiedCount++
}

func (r *Restaurant) ReleaseTable() error {
	r.cond.L.Lock()
	defer r.cond.L.Unlock()
	if r.occupiedCount == 0 {
		return errors.New("cant free more tables")
	}

	r.occupiedCount--

	r.cond.Signal()
	return nil
}

func (r *Restaurant) IsFree() bool {
	return r.occupiedCount < r.tableCount
}
