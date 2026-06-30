package t8

import "sync"

func Parse(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		for value := range in {
			parsed := "parsed - " + value
			// fmt.Println(parsed)
			out <- parsed
		}
		close(out)
	}()

	return out
}

func Split(in <-chan string, n int) []chan string {
	out := make([]chan string, 0, n)
	for range n {
		out = append(out, make(chan string))
	}
	go func() {
		for _, outChan := range out {
			go func() {
				for val := range in {
					outChan <- val
				}
				close(outChan)
			}()
		}
	}()
	return out
}

func Send(in []chan string) <-chan string {
	out := make(chan string)
	var wg sync.WaitGroup
	for _, inChan := range in {
		wg.Go(func() {
			for val := range inChan {
				sentData := "sent - " + val
				out <- sentData
			}
		})
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
