package main

import "fmt"

func main() {
	slice := make([]string, 0, 5)
	slice = append(slice, "0", "1", "2", "3")
	fmt.Println(slice, len(slice), cap(slice)) // [0 1 2 3] 4 5
	addToSlice1(slice)
	fmt.Println(slice, len(slice), cap(slice)) // [0 1 2 one] 4 5
	addToSlice2(slice)
	fmt.Println(slice, len(slice), cap(slice)) //
}

func addToSlice1(slice []string) {
	// Значение slice из main поменялось, так как ипользовался тот же
	// массив, и он не перекопировался
	slice = append(slice[1:3], "one")
}

func addToSlice2(slice []string) {
	// здесь использовалась копия массива
	slice = append(slice, "two")
}
