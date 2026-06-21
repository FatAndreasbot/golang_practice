package stack

type stack[T any] struct {
	slice []T
	top   int
}

func (s *stack[T]) Push(v T) {
	s.slice = append(s.slice, v)
	s.top += 1
}

func (s *stack[T]) Pop() T {
	result := s.slice[s.top-1]
	s.top -= 1
	return result
}

func New[T any]() *stack[T] {
	return &stack[T]{
		slice: make([]T, 0),
		top:   0,
	}
}
