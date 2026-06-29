package t6

import (
	"errors"
	"fmt"
	"time"
)

func DBReplica(name string, in <-chan int) {
	for data := range in {
		fmt.Printf("Запись в %s: %d\n", name, data)
		time.Sleep(100 * time.Millisecond) // Имитация задержки записи
	}
	fmt.Printf("Реплика %s закрыта\n", name)
}

func TeeChannel[T any](in <-chan T, count int) ([]<-chan T, error) {
	if count < 2 {
		return nil, errors.New("need at least 2 output channels")
	}

	ret := make([]<-chan T, 0, count)

	go func() {
		for value := range in {

			for range count {
				out := make(chan T)
				out <- value
				ret = append(ret, out)

			}
		}
	}()

	return ret, nil
}
