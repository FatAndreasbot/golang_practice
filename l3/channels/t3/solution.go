package main

import "fmt"

func main() {
	naturals := make(chan int, 65536)
	squares := make(chan int, 65536)

	go func() {
		for natural := range naturals {
			squares <- (natural * natural)
		}
	}()

	for i := range 50 {
		naturals <- i
	}

	for range 50 {
		square := <-squares
		fmt.Printf("%d ", square)
	}
}
