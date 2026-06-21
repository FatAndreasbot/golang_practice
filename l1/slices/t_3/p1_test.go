package t3_test

import (
	"fmt"
	"t3"
	"testing"
)

func TestRemoveUnordered(t *testing.T) {
	s := []int{0, 1, 2, 3, 4}
	fmt.Println("\nRemove unordered")

	fmt.Println(s)
	s = t3.RemoveUnordered(s, 2)
	fmt.Println(s)
}

func TestRemoveOrdered(t *testing.T) {
	s := []int{0, 1, 2, 3, 4}
	fmt.Println("\nRemove ordered")

	fmt.Println(s)
	s = t3.RemoveOrdered(s, 2)
	fmt.Println(s)
}
