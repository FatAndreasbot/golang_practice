```go
package main

import "fmt"

func main() {
	first := []int{10, 20, 30, 40}     // создаем массив чисел
	second := make([]*int, len(first)) // создаем массив указателей на числа
	for i := range first {
		second[i] = &first[i] // в массив указателей кладем указатели на числа первого массива
	}

	first[0] = 1 // для проверки - заменяю первый элемент массива
	fmt.Println(*second[0], *second[1])
}
