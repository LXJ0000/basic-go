package list

import "github.com/LXJ0000/basic-go/lib/errs"

type node struct {
	val  int
	prev *node
	next *node
}
type LinkedList struct {
	head, tail *node
	size       int
}

func NewLinkList(nums ...int) *LinkedList {
	head := &node{}
	tail := &node{}
	head.next = tail
	tail.prev = head

	l := &LinkedList{
		head: head,
		tail: tail,
	}

	for _, num := range nums {
		l.PushBack(num)
	}

	return l
}

// PushFront 在首部追加元素
func (l *LinkedList) PushFront(num int) {
	midNode := &node{
		val: num,
	}
	midNode.prev = l.head
	midNode.next = l.head.next
	l.head.next.prev = midNode
	l.head.next = midNode
	l.size++
}

// PushBack 在末尾追加元素
func (l *LinkedList) PushBack(num int) {
	midNode := &node{
		val: num,
	}
	midNode.prev = l.tail.prev
	midNode.next = l.tail
	l.tail.prev.next = midNode
	l.tail.prev = midNode
	l.size++
}

// PopFront 删除首部元素
func (l *LinkedList) PopFront() {
	midNode := l.head.next
	midNode.prev.next = midNode.next
	midNode.next.prev = midNode.prev
	l.size--
}

// PopBack 删除末尾元素
func (l *LinkedList) PopBack() {
	midNode := l.tail.prev
	midNode.prev.next = midNode.next
	midNode.next.prev = midNode.prev
	l.size--
}

func (l *LinkedList) Size() int {
	return l.size
}

// ToSlice 转化为 Slice 并返回
func (l *LinkedList) ToSlice() []int {
	nums := make([]int, 0, l.size)
	h := l.head.next
	for h != l.tail {
		val := h.val
		nums = append(nums, val)
		h = h.next
	}
	return nums
}

// Add 在指定位置添加一个数，index <= 0 则添加至首部; index >= size 则添加至末尾
func (l *LinkedList) Add(index int, num int) {
	if index <= 0 {
		l.PushFront(num)
		return
	}
	if index >= l.size {
		l.PushBack(num)
		return
	}
	h := l.head
	for i := 0; i < index; i++ {
		h = h.next
	}
	node := &node{val: num}
	node.prev = h
	node.next = h.next
	h.next.prev = node
	h.next = node
}

// Delete 删除指定位置的数
func (l *LinkedList) Delete(index int) (int, error) {
	if index < 0 || index >= l.size {
		return 0, errs.NewErrIndexOutOfRange(l.size, index)
	}
	midNode := l.head.next
	for i := 0; i < index; i++ {
		midNode = midNode.next
	}
	midNode.prev.next = midNode.next
	midNode.next.prev = midNode.prev
	l.size--
	return midNode.val, nil
}
