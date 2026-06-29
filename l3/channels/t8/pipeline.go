package t8

func Parse(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		value := <-in
		out <- "parsed - " + value
	}()

	return out
}

func Split(in chan string, n int) []chan string {
	out := make([]chan string, 0, n)
	for range n {
		out = append(out, make(chan string))
	}

	go func() {
		for value := range in {
			for range n {
				if len(out[n]) == cap(out[n]) {
					continue
				}
				out[n] <- value
			}
		}
	}()

	return out
}
