package linked_list

type LinkedList interface {
	AddFirst(data interface{})
	AddLast(data interface{})
	Remove(data interface{})
}