package main

import (
	"fmt"
	"strings"
)

// TODO перевести слова в н рег
// WordFrequency принимает строку текста и возвращает map с частотой слов.
func WordFrequency(text string) map[string]int {
	wordFreq := make(map[string]int)

	for _, val := range strings.Split(text, " ") {
		_, exists := wordFreq[val]
		if exists {
			wordFreq[val] += 1
		} else {
			wordFreq[val] = 1
		}
	}

	return wordFreq
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
