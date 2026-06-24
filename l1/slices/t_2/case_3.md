```go
package main

import "fmt"

func test(testSlice []string) {
	testSlice[len(testSlice)-1] = "Пока"
}

func main() {
	testSlice := make([]string, 0, 3)
	testSlice = append(testSlice, "Привет")
	testSlice = append(testSlice, "Привет")

	testSlice = append(testSlice, "") // вместо указателей я увеличил размер слайса
	// предварительно, а заетм просто заенил значение в элементе массива
	test(testSlice)

	fmt.Println(testSlice)
}
