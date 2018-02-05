package stack

import (
	"errors"
)

type StackArray struct {
	items []string
	size uint64
	top uint64
}

func NewStackArray(size uint64) *StackArray {
	return &StackArray{
		items: make([]string, size),
		size: size,
	}
}

// Push pushes data to the top of the Stack
func (s *StackArray) Push(data interface{}) error {

	// if we don't have empty slots anymore
	if s.top == s.size {
		return errors.New("stack overflow")
	}

	// assign data to the top
	s.items[s.top] = data.(string)

	// increase top
	s.top += 1

	return nil
}

// Pop returns data on the head of the Stack(removing head)
func (s *StackArray) Pop() (interface{}, error) {
	// if stack is empty return error
	if s.IsEmpty() {
		return createStackIsEmptyError()
	}

	// decrease top
	s.top -= 1

	// save the top value to the local variable
	data := s.items[s.top]

	return data, nil
}

// Peek returns data on the head of the Stack(not removing head)
func (s *StackArray) Peek() (interface{}, error) {

	// if stack is empty return error
	if s.IsEmpty() {
		return createStackIsEmptyError()
	}

	return s.items[s.top - 1], nil
}

// IsEmpty returns whether Stack is empty or not
func (s *StackArray) IsEmpty() bool {

	return s.top == 0
}

// Size returns size of the Stack
func (s *StackArray) Size() uint64 {
	return s.top
}
