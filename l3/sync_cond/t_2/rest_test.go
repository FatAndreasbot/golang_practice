package t2_test

import (
	"t2"
	"testing"
	"time"
)

func TestReastaurant(t *testing.T) {
	rest := t2.NewReastaurant(5)

	go func() {
		for range 6 {
			rest.OccupyTable()
		}
	}()

	time.Sleep(time.Millisecond * 300)
	if rest.IsFree() {
		t.Error("the restaurant should be full")
	}
	_ = rest.ReleaseTable()
	time.Sleep(time.Millisecond * 300)

	if rest.IsFree() {
		t.Error("the restaurant should be full")
	}

	_ = rest.ReleaseTable()
	_ = rest.IsFree()
	time.Sleep(time.Millisecond * 300)

	if !rest.IsFree() {
		t.Error("the restaurant should have a free table")
	}

	for range 4 {
		_ = rest.ReleaseTable()
	}

	err := rest.ReleaseTable()
	if err == nil {
		t.Error("i should get an error")
	}
}
