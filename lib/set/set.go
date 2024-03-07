package set

type Set[T comparable] interface {
	Add(key T)
	Delete(key T)
	Exists(key T) bool
}

type MapSet[T comparable] struct {
	m map[T]struct{}
}

func NewMapSet[T comparable](size int) *MapSet[T] {
	return &MapSet[T]{
		m: make(map[T]struct{}, size),
	}
}

func (s *MapSet[T]) Add(key T) {
	s.m[key] = struct{}{}
}

func (s *MapSet[T]) Delete(key T) {
	delete(s.m, key)
}

func (s *MapSet[T]) Exists(key T) bool {
	_, ok := s.m[key]
	return ok
}
