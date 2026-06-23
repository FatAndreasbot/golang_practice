package t3

import (
	"errors"
	"fmt"
	"math/rand"
)

var (
	ErrNotFound  = errors.New("ресурс не найден")
	TimeoutError = errors.New("таймаут операции")
)

func ProcessError(err error) {
	if errors.Is(err, TimeoutError) {
		fmt.Println("Требуется повторная попытка")
	}
	if errors.Is(err, ErrNotFound) {
		fmt.Println("Требуется повторная попытка")
	}
}

func SimulateRequest() error {
	if rand.Intn(2) == 0 {
		return fmt.Errorf("запрос не выполнен: %w", TimeoutError)
	}
	if rand.Intn(10) < 3 {
		return fmt.Errorf("ошибка: %w", ErrNotFound)
	}
	if rand.Intn(5) == 0 {
		return errors.New("неизвестная ошибка")
	}
	return nil
}
