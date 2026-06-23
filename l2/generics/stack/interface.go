package stack

type Stacker[T any] interface {
	Push(value T)
	Pop() (T, bool)
	Peek() (T, bool)
	IsEmpty() bool
}
