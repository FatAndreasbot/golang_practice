package t5_test

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"t5"
	"testing"
)

const WordCount = 458 // https://www.lipsum.com/feed/html (5 paragraphs)

type Counts struct {
	WordCount int
	FileName  string
	Error     error
}

func TestFanout(t *testing.T) {
	contentDir := "./files"
	files, err := os.ReadDir(contentDir)
	if err != nil {
		t.Error(err)
	}
	fileCount := len(files)
	input := make(chan string, fileCount)
	var wg sync.WaitGroup

	outputs := t5.FanOut(input, fileCount, func(filename string) Counts {
		defer wg.Done()
		wg.Add(1)
		fileContent, err := t5.ReadFile(filename)
		if err != nil {
			return Counts{
				WordCount: 0,
				FileName:  filename,
				Error: errors.Join(
					err,
					fmt.Errorf("got error on %s", filename),
				),
			}
		}

		wordCount := t5.CountWords(fileContent)

		return Counts{
			WordCount: wordCount,
			FileName:  filename,
			Error:     nil,
		}
	})

	for _, file := range files {
		fileData, err := file.Info()
		if err != nil {
			continue
		}
		input <- filepath.Join(contentDir, fileData.Name())
	}

	close(input)
	wg.Wait()

	var allWordCount int
	for _, output := range outputs {
		data := <-output

		allWordCount += data.WordCount
	}

	if allWordCount != WordCount {
		t.Errorf("the ammount of words is wrong\ngot %d, expected %d", allWordCount, WordCount)
	}

	fmt.Println(files)
}
