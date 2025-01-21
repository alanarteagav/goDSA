package dataStructures

// Queue struct for generic types
type Queue[T any] struct {
	items []T
}

// Enqueues an item
func (s *Queue[T]) Enqueue(item T) {
	s.items = append(s.items, item)
}

// Dequeues an item
func (s *Queue[T]) Dequeue() T {
	item := s.items[0]
	s.items = s.items[1:]
	return item
}

// Checks if the queue is empty
func (s *Queue[T]) IsEmpty() bool {
	return len(s.items) == 0
}

/* The following metods are used to implement the Collection interface */

// Adds an item
func (s *Queue[T]) Add(item T) {
	s.Enqueue(item)
}

// Removes an item
func (s *Queue[T]) Remove() T {
	return s.Dequeue()
}
