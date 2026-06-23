package t2

import (
	"fmt"
)

func FormattedError(age int) error {
	return fmt.Errorf("ошибка: возраст %d недопустим", age)
}

func WrapError(err error, msg string) error {
	return fmt.Errorf("%s : %w", msg, err)
}
