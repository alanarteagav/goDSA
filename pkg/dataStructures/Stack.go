package dataStructures

type Stack[T Ordered] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() T {
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

/* The following metods are used to implement the Collection interface */
func (s *Stack[T]) Add(item T) {
	s.Push(item)
}

func (s *Stack[T]) Remove() T {
	return s.Pop()
}
