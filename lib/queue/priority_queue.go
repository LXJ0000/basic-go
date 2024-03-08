package queue

import "github.com/LXJ0000/basic-go/lib"

type PriorityQueue[T lib.Ordered] struct {
	data []T
	size int
}

func NewPriorityQueue[T lib.Ordered]() *PriorityQueue[T] {
	return &PriorityQueue[T]{
		data: []T{0},
		size: 0,
	}
}

func (h *PriorityQueue[T]) down(u int) {
	t, n := u, h.size

	if u*2 <= n && h.data[t] > h.data[u*2] {
		t = u * 2
	}
	if u*2+1 <= n && h.data[t] > h.data[u*2+1] {
		t = u*2 + 1
	}
	if t != u {
		h.data[u], h.data[t] = h.data[t], h.data[u]
		h.down(t)
	}
}

func (h *PriorityQueue[T]) up(u int) {
	for u/2 != 0 && h.data[u] < h.data[u/2] {
		h.data[u], h.data[u/2] = h.data[u/2], h.data[u]
		u >>= 1
	}
}

func (h *PriorityQueue[T]) Push(x T) {
	h.data = append(h.data, x)
	h.size++
	h.up(h.size)
}

func (h *PriorityQueue[T]) Pop() T {
	move := h.data[1]
	h.data[1], h.data[h.size] = h.data[h.size], h.data[1]
	h.size--
	h.down(1)
	return move
}

func (h *PriorityQueue[T]) Front() T {
	return h.data[1]
}

func (g *PriorityQueue[T]) Size() int {
	return g.size
}
func (g *PriorityQueue[T]) Empty() bool {
	return g.size == 0
}
