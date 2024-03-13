package list

import (
	"fmt"
	"testing"
)

func TestLinkedList_PushFront(t *testing.T) {
	l := NewLinkList(0, 0)
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)
	l.PushBack(5)
	l.PushFront(0)
	l.PushFront(-1)
	l.PushFront(-2)
	l.PushFront(-3)
	//[-3 -2 -1 0 1 2 3 4 5]
	l.PopFront()
	l.PopFront()
	l.PopBack()
	l.PopFront()
	fmt.Println(l.ToSlice())

	l.Add(-1, 11)
	l.Add(99, 99)
	l.Add(3, 33)
	fmt.Println(l.ToSlice())
}
