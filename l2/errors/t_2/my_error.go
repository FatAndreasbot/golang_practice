package t2

import "fmt"

type MyError struct {
	Msg  string
	Code int
}

func (err MyError) Error() string {
	return fmt.Sprintf("%s: Код ошибки - %d", err.Msg, err.Code)
}

func StructError() error {
	return MyError{
		Msg:  "не найдено",
		Code: 404,
	}
}
