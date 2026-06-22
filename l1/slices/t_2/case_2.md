package main

import (
	"fmt"
	"strings"
)

func chengeSlice(arr []string) {
	arr[0] = "Goodbye"
}

func appendSomeData(arr *[]string) {
	*arr = append(*arr, "!")
}

func main() {
	someSlice := []string{"Hello", "World"}
	chengeSlice(someSlice)
	appendSomeData(&someSlice) // вместо того, чтобы отправлять слайс,
	// использую указатель на слайс

	fmt.Println(strings.Join(someSlice, ""))
}
