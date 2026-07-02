package main

import (
	"bytes"
	"fmt"
	"sync"
)

const CAP_SYMBOL_DIFFERENCE byte = 'a' - 'A'

var alphabetArray []byte = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

var stringBufferPool sync.Pool = sync.Pool{
	New: func() any { return bytes.NewBuffer([]byte{}) },
}

// не совсем понял как и зачем тут применять pool, поэтому вот такое написал
// UPD до меня дошло!!! в строке 20 я аллоцирую новый массив. а надо брать инстанс из pool
func ProcessString(s string) string {
	buffer := stringBufferPool.Get().(*bytes.Buffer)
	defer stringBufferPool.Put(buffer)
	defer buffer.Reset()
	for i := range s {
		char := s[i]
		if char < 'a' || char > 'z' {
			buffer.WriteByte(char)
			continue
		}
		buffer.WriteByte(alphabetArray[char-'a'])
	}

	return buffer.String()
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
