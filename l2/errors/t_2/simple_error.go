package t2

import "errors"

func SimpleError() error {
	return errors.New("простая ошибка")
}
