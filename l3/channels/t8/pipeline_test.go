package t8_test

import (
	"fmt"
	"strings"
	"sync"
	"t8"
	"testing"
)

func TestPipeline(t *testing.T) {
	inputStringCount := 10
	splits := 5

	input := make(chan string)
	output := t8.Send(t8.Split(t8.Parse(input), splits))
	var wg sync.WaitGroup

	for i := range inputStringCount {
		wg.Go(func() {
			input <- strings.Repeat("*", i+1)
		})
	}

	go func() {
		wg.Wait()
		close(input)
	}()
	var fullOutput string
	for value := range output {
		// fmt.Println(value)
		fullOutput += value + "\n"
	}

	fmt.Println(fullOutput)

}
