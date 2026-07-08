package t5

import (
	"errors"
	"os"
	"strings"
)

func FanOut[InputT any, ReturnT any](chanInput <-chan InputT, count int, work func(InputT) ReturnT) []<-chan ReturnT {
	ret := make([]<-chan ReturnT, 0, count)
	// TODO close channel
	for range count {
		output := make(chan ReturnT)
		ret = append(ret, output)
		go func() {
			for input := range chanInput {
				output <- work(input)
			}

		}()
	}

	return ret
}

func ReadFile(filename string) (string, error) {
	rawContent, err := os.ReadFile(filename)
	if err != nil {
		return "", errors.Join(
			err,
			errors.New("could not read file"),
		)
	}
	content := string(rawContent)

	return content, nil
}

func CountWords(data string) int {
	allWords := strings.Split(data, " ")

	return len(allWords)
}
