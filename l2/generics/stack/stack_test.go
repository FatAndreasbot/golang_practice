package stack_test

import (
	"stack"
	"testing"
)

func TestStack_PushPop(t *testing.T) {
	s := stack.NewStack[int]()

	s.Push(1)
	s.Push(2)
	s.Push(3)

	tests := []struct {
		expected int
	}{
		{3},
		{2},
		{1},
	}

	for range 2 {
		got, _ := s.Peek()
		if got != 3 {
			t.Errorf("Peek() = %d; ожидалось %d", got, 3)
		}
	}

	for _, tc := range tests {
		got, _ := s.Pop()
		if got != tc.expected {
			t.Errorf("Pop() = %d; ожидалось %d", got, tc.expected)
		}
	}

	_, isNotEmpty := s.Peek()
	if isNotEmpty {
		t.Error("the stack should be empty at this point")
	}
}
