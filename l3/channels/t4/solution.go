package t4

import "sync"

func MergeChannels[T any](channels ...<-chan T) <-chan T {
	ret := make(chan T)
	var wg sync.WaitGroup
	wg.Add(len(channels))

	for _, channel := range channels {
		go func() {
			defer wg.Done()
			for value := range channel {
				ret <- value
			}
		}()
	}

	go func() {
		wg.Wait()
		close(ret)
	}()

	return ret
}
