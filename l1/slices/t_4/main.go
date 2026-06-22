package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// 1
	first := []int{1, 2, 3, 4, 5}
	first = nil
	fmt.Println("first:", first, ":", len(first), ":", cap(first))
	// first: []:0:0

	// 2
	second := []int{1, 2, 3, 4, 5}
	second = second[:0]
	fmt.Println("second:", second, ":", len(second), ":", cap(second))
	// second:[1]:1:5

	// 3
	third := []int{1, 2, 3, 4, 5}
	clear(third)
	fmt.Println("third:", third, ":", len(third), ":", cap(third))
	// third: [0,0,0,0,0]:5:5

	// 4
	fourth := []int{1, 2, 3, 4, 5}
	clear(fourth[1:3])
	fmt.Println("fourth:", fourth, ":", len(fourth), ":", cap(fourth))
	// fourth:[1 0 0 0 5]:5:5

	// 5
	slice := make([]int, 3, 6)
	array := [3]int(slice[:3])
	slice[0] = 10

	fmt.Println("slice = ", slice, len(slice), cap(slice))
	fmt.Println("array =", array, len(array), cap(array))
	// slice = [10 0 0] 3 6
	// array = [0 0 0] 3 3

	// 6 В каких случаях Slice пустой или нулевой
	// 1
	var data []string
	fmt.Println("var data []string:") // nil pointer at array
	fmt.Printf("\tempty=%t nil=%t size=%d data=%p\n", len(data) == 0, data == nil, unsafe.Sizeof(data), unsafe.SliceData(data))
	// 2
	data = []string(nil) // nil pointer at array
	fmt.Println("data = []string(nil):")
	fmt.Printf("\tempty=%t nil=%t size=%d data=%p\n", len(data) == 0, data == nil, unsafe.Sizeof(data), unsafe.SliceData(data))
	// 3
	data = []string{} // empty array
	fmt.Println("data = []string{}")
	fmt.Printf("\tempty=%t nil=%t size=%d data=%p\n", len(data) == 0, data == nil, unsafe.Sizeof(data), unsafe.SliceData(data))
	// 4
	data = make([]string, 0) // empty array
	fmt.Println("data =make([]string,0)")
	fmt.Printf("\tempty=%t nil=%t size=%d data=%p\n", len(data) == 0, data == nil, unsafe.Sizeof(data), unsafe.SliceData(data))

	empty := struct{}{}
	fmt.Println("empty struct address ", unsafe.Pointer(&empty))
}
