package dataStructures

type Queue[T Ordered] struct {
	items []T
}

func (s *Queue[T]) Enqueue(item T) {
	s.items = append(s.items, item)
}

func (s *Queue[T]) Dequeue() T {
	item := s.items[0]
	s.items = s.items[1:]
	return item
}

func (s *Queue[T]) IsEmpty() bool {
	return len(s.items) == 0
}

/* The following metods are used to implement the Collection interface */
func (s *Queue[T]) Add(item T) {
	s.Enqueue(item)
}

func (s *Queue[T]) Remove() T {
	return s.Dequeue()
}
