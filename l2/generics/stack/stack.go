package stack

type stack[T any] struct {
	slice []T
}

func (s *stack[T]) Push(v T) {
	s.slice = append(s.slice, v)
}

func (s *stack[T]) Peek() (T, bool) {
	if len(s.slice) == 0 {
		var zero T
		return zero, false
	}

	result := s.slice[len(s.slice)-1]
	return result, true
}

func NewStack[T any]() *stack[T] {
	return &stack[T]{
		slice: make([]T, 0),
	}
}

func (s *stack[T]) Pop() (T, bool) {
	val, isNotEmpty := s.Peek()
	if isNotEmpty {
		s.slice = s.slice[:len(s.slice)-1]
	}

	return val, isNotEmpty
}

func (s stack[T]) IsEmpty() bool {
	return len(s.slice) == 0
}
