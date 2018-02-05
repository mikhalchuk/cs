package stack

import "errors"

type Stack interface {
	Push(data interface{}) error
	Pop() (interface{}, error)
	Peek() (interface{}, error)
	IsEmpty() bool
	Size() uint64
}

// createStackIsEmptyError creates error in case of empty Stack
func createStackIsEmptyError() (string, error) {
	return "", errors.New("stack is empty")
}
