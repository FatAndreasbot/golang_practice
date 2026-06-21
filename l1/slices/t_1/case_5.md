package main

import "fmt"

func main() {
	arr := []int{1, 2, 3}
	src := arr[:1]
	foo(src)
	fmt.Println(src) // [1]
	fmt.Println(arr) // [1 5 3]
}

func foo(src []int) {
	// src - копия из main, но указывают на один и тот же
	// массив. таким образом в arr заменяется элемент,
	// однако в main.src не меняется capacity
	src = append(src, 5)
}
