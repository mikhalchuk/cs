package stack

type NodeList struct {
	data interface{}
	next *NodeList
}

type StackList struct {
	head *NodeList
	size uint64
}

// Push pushes data to the top of the Stack
func (s *StackList) Push(data interface{}) error {

	// create new node
	newHead := &NodeList{
		data: data,
		next: nil,
	}

	// if stack is not empty, assign to the node's next the current head
	if s.head != nil {
		newHead.next = &NodeList{}
		*newHead.next = *s.head
	}

	// assign stack's head as a new node
	s.head = newHead

	// increase size
	s.size += 1

	return nil
}

// Pop returns data on the head of the Stack(removing head)
func (s *StackList) Pop() (interface{}, error) {

	// if stack is empty return error
	if s.IsEmpty() {
		return createStackIsEmptyError()
	}

	// save head data in a local variable
	data := s.head.data

	// change pointer of the stack to the next one
	if s.head.next != nil {
		*s.head = *s.head.next
	} else {
		s.head = nil
	}

	// decrease size
	s.size -= 1

	return data, nil
}

// Peek returns data on the head of the Stack(not removing head)
func (s *StackList) Peek() (interface{}, error) {

	// if stack is empty return error
	if s.IsEmpty() {
		return createStackIsEmptyError()
	}

	return s.head.data, nil
}

// IsEmpty returns whether Stack is empty or not
func (s *StackList) IsEmpty() bool {
	return s.size == 0
}

// Size returns size of the Stack
func (s *StackList) Size() uint64 {
	return s.size
}
