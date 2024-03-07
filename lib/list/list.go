package list

type node[T any] struct {
	val  T
	prev *node[T]
	next *node[T]
}

type LinkedList[T any] struct {
	head *node[T]
	tail *node[T]
	size int
}

func NewLinkedList[T any]() *LinkedList[T] {
	head, tail := &node[T]{}, &node[T]{}
	head.next = tail
	tail.prev = head
	return &LinkedList[T]{
		head: head,
		tail: tail,
	}
}
