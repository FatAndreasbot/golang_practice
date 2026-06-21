package main

import (
	"fmt"
)

func main() {
	nums := make([]int, 1, 3)    // array - [0 0 0]
	fmt.Println(nums)            // [0] - size = 1
	appendSlice(nums, 1)         // array - [0 1 0]
	fmt.Println(nums)            // [0] - size = 1
	copySlice(nums, []int{2, 3}) // array [2 3 0]
	fmt.Println(nums)            // [2] - size = 1
	mutateSlice(nums, 1, 4)      // panic
	fmt.Println(nums)            // ---
}

func appendSlice(sl []int, val int) {
	sl = append(sl, val)
}

func copySlice(sl, cp []int) {
	copy(sl, cp)
}

func mutateSlice(sl []int, idx, val int) {
	sl[idx] = val
}
