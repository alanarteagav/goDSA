package dataStructures

// Stack struct for generic types
type Stack[T any] struct {
	items []T
}

// Pushes an item
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pops an item
func (s *Stack[T]) Pop() T {
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

// Checks if the stack is empty
func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

/* The following metods are used to implement the Collection interface */

// Adds an item
func (s *Stack[T]) Add(item T) {
	s.Push(item)
}

// Removes an item
func (s *Stack[T]) Remove() T {
	return s.Pop()
}
