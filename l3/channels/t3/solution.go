package main

import (
	"fmt"
	"sync"
)

// TODO check
func main() {
	naturals := make(chan int)
	squares := make(chan int)

	var wg sync.WaitGroup

	go func() {
		for natural := range naturals {
			squares <- (natural * natural)
		}
	}()

	go func() {
		for i := range 50 {
			wg.Add(1)
			naturals <- i
		}
	}()

	go func() {
		for range 50 {
			square := <-squares
			fmt.Printf("%d ", square)
		}
	}()

	wg.Wait()
	close(naturals)
	close(squares)
}
