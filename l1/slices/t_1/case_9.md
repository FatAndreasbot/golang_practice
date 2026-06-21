package main

import (
	"fmt"
)

func main() {
	slice := make([]int, 3, 4) // array - [0 0 0 0] size - 3
	appendingSlice(slice[:1])  // array - [0 1 0 0] size - 3
	fmt.Println(slice)         // [0 1 0]
}

func appendingSlice(slice []int) {
	slice = append(slice, 1)
}
