package t3_test

import (
	"slices"
	"t3"
	"testing"
)

func TestRemoveAllByValue(t *testing.T) {
	s := []int{0, 1, 2, 2, 2, 3, 4}

	s2 := t3.RemoveAllByValue(s, 2)
	idealSlice := []int{0, 1, 3, 4}

	if slices.Compare(s2, idealSlice) != 0 {
		t.Error("got wrong slice")
	}
}
