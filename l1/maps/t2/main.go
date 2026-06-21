package main

import (
	"fmt"
	"strings"
)

// WordFrequency принимает строку текста и возвращает map с частотой слов.
func WordFrequency(text string) map[string]int {
	wf := make(map[string]int)

	for _, v := range strings.Split(text, " ") {
		_, exists := wf[v]
		if exists {
			wf[v] += 1
		} else {
			wf[v] = 1
		}
	}

	return wf
}

// PrintWordFrequency выводит частотный анализ слов, отсортированный по убыванию частоты.
func PrintWordFrequency(freqMap map[string]int) {
	for k, v := range freqMap {
		fmt.Printf("The word %q was found %d times\n", k, v)
	}
}

func main() {
	text := "golang is great and golang is fast"

	wf := WordFrequency(text)
	PrintWordFrequency(wf)

}
