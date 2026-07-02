package main

import (
	"fmt"
	"sync"
)

const CAP_SYMBOL_DIFFERENCE byte = 'a' - 'A'

var stringBufferPool sync.Pool = sync.Pool{
	New: func() any { return []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ") },
}

// не совсем понял как и зачем тут применять pool, поэтому вот такое написал
func ProcessString(s string) string {
	alphabetArray := stringBufferPool.Get().([]byte)
	buffer := []byte(s)
	for i := range buffer {
		char := buffer[i]
		if char < 'a' || char > 'z' {
			continue
		}
		buffer[i] = alphabetArray[char-'a']
	}

	stringBufferPool.Put(alphabetArray)

	return string(buffer)
}

func main() {
	examples := []string{
		"hello, world!",
		"gopher",
		"lorem ipsum dolor sit amet",
	}

	for _, s := range examples {
		processed := ProcessString(s)
		fmt.Printf("Original: %q\nProcessed: %q\n\n", s, processed)
	}
}
