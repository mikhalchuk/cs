package stack

import (
	"testing"
)

func run(f func(stack Stack)) {

	for _, stack := range []Stack{&StackList{}, NewStackArray(10)} {
		f(stack)
	}
}

func TestStack_IsEmpty(t *testing.T) {
	f := func(stack Stack) {
		if !stack.IsEmpty() {
			t.Error("Newly created stack should be empty")
		}
	}

	run(f)
}

func TestStack_IsNotEmpty(t *testing.T) {
	f := func(stack Stack) {
		_ = stack.Push("a")

		if stack.IsEmpty() {
			t.Error("Stack should not be empty")
		}
	}
	run(f)
}

func TestStack_SizeIsZero(t *testing.T) {
	f := func(stack Stack) {
		if stack.Size() != 0 {
			t.Error("Newly created stack should have size 0")
		}
	}
	run(f)
}

func TestStack_SizeIsNotZero(t *testing.T) {
	f := func(stack Stack) {
		_ = stack.Push("a")
		if stack.Size() <= 0 {
			t.Error("Stack should have size 0")
		}
	}

	run(f)
}

func TestStack_PopWhenStackIsEmpty(t *testing.T) {
	f := func(stack Stack) {
		_, err := stack.Pop()
		if err == nil {
			t.Error("Must be an error when stack is empty")
		}
	}

	run(f)
}

func TestStack_FIFO(t *testing.T) {
	f := func(stack Stack) {
		items := []string{"a", "b", "c", "d", "f", "e"}
		for _, c := range items {
			_ = stack.Push(c)
		}

		for i := len(items) - 1; i >= 0; i-- {
			item, _ := stack.Pop()
			if items[i] != item.(string) {
				t.Errorf("Order is wrong: should be %s received %s", items[i], item)
			}
		}

		stack.Push("sasad")
		stack.Pop()
		stack.Peek()
		stack.Push("333")
		stack.Push("45555")
		stack.Pop()
		stack.Pop()

		_, err := stack.Pop()
		if err == nil {
			t.Error("Must be an error when stack is empty")
		}
	}

	run(f)
}

func TestStack_PeekStackIsEmpty(t *testing.T) {
	f := func(stack Stack) {
		_, err := stack.Peek()
		if err == nil {
			t.Error("Must be an error when stack is empty")
		}
	}

	run(f)
}

func TestStack_PeekStackIsNotEmpty(t *testing.T) {
	f := func(stack Stack) {
		data := "a"
		_ = stack.Push(data)

		item, _ := stack.Peek()

		if data != item {
			t.Error("They should be equal")
		}

		if stack.Size() != 1 {
			t.Error("Peek should not remove node")
		}
	}

	run(f)
}

func TestStackArray_PushOverflow(t *testing.T) {
	stack := NewStackArray(1)
	_ = stack.Push("a")

	err := stack.Push("b")
	if err == nil {
		t.Error("Should have been received overflow error")
	}
}