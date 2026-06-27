package t3_test

import (
	"slices"
	"t3"
	"testing"
)

func TestRemoceDuplicates(t *testing.T) {
	s := []int{1, 2, 2, 2, 3, 4}
	s2 := []int{1, 2, 3, 4}

	if slices.Compare(t3.RemoveDuplicates(s), s2) != 0 {
		t.Error("Remove duplicates does not work")
	}
}

func TestReomveIf(t *testing.T) {
	predicate := func(v int) bool {
		return v%2 == 1
	}
	s := []int{1, 2, 3, 4, 5, 6}
	s2 := []int{2, 4, 6}

	res := t3.RemoveIf(s, predicate)
	if slices.Compare(res, s2) != 0 {
		t.Error("Filter does not work")
	}
}
