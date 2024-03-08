package queue

import (
	"fmt"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	q := NewPriorityQueue[float64]()
	q.Push(1.1)
	q.Push(2)
	q.Push(-1.2)
	q.Push(4.999)

	for !q.Empty() {
		fmt.Println(q.Pop())
	}

}
