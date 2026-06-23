package main

import (
	"fmt"
)

type MyError struct {
	err_message string
}

func (m MyError) Error() string {
	return m.err_message
}

func NewErr(msg string) MyError {
	return MyError{
		err_message: msg,
	}
}

func handle() error {
	return NewErr("New error")
}

func main() {

	err := handle()

	if err == nil {
		panic("expected an error message")
	}
	fmt.Println(err)
}
