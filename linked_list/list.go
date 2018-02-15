package linked_list

type Node struct {
	data interface{}
	next *Node
}

type List struct {
	head *Node
}

func (l *List) AddFirst(data interface{}) {
	// create new node
	newHead := &Node{
		data: data,
		next: nil,
	}

	// if list is not empty, assign to the node's next the current head
	if l.head != nil {
		newHead.next = &Node{}
		*newHead.next = *l.head
	}

	// assign list's head as a new node
	l.head = newHead
}

func (l *List) AddLast(data interface{}) {
	// if list is empty insert as a first node
	if l.head == nil {
		l.AddFirst(data)
		return
	}

	// find last node
	tmp := l.getLastNode()

	// assign new node as a last one
	tmp.next =  &Node{
		data: data,
		next: nil,
	}
}

func (l *List) Remove(data interface{}) {

}

func (l *List) ToArray() []interface{} {
	list := make([]interface{}, 0)

	tmp := l.head
	for tmp != nil {
		list = append(list, tmp.data)
		tmp = tmp.next
	}

	return list
}

func (l *List) getLastNode() *Node {
	tmp := l.head
	for tmp.next != nil {
		tmp = tmp.next
	}

	return tmp
}