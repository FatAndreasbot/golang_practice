package stack

type Stacker[T any] interface {
	Push(v T)
	Pop() T
}
