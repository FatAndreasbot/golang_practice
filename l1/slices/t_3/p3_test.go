package t3

import (
	"fmt"
	"testing"
)

func TestReducingCapacity(t *testing.T) {
	s := make([]int, 0, 20)
	s = append(s, 0, 1, 2, 3, 4, 5, 6, 7)
	oldcap := cap(s)
	s = reduceCapacity(s)
	newcap := cap(s)

	fmt.Printf("\nold capacity: %d. new capacity: %d\n", oldcap, newcap)
	if oldcap <= newcap {
		t.Error("capacity was not decreased")
	}

}
