package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	cnt := 100
	for i := range cnt {
		wg.Go(func() {
			fmt.Println(i)
		})
	}

	wg.Wait()
}
