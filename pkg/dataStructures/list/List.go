package list

import "errors"

// Doubly linked list node
type Node[T comparable] struct {
	value    T
	previous *Node[T]
	next     *Node[T]
}

// Doubly Linked List struct
type List[T comparable] struct {
	head   *Node[T]
	tail   *Node[T]
	length int
}

// List Factory Function
func NewList[T comparable]() *List[T] {
	lp := &List[T]{
		head: nil,
		tail: nil,
	}
	return lp
}

// Adds to the end of the list
func (l *List[T]) Add(item T) {
	if l.head == nil {
		newNode := &Node[T]{
			value:    item,
			previous: nil,
			next:     nil,
		}
		l.head = newNode
		l.tail = newNode
		l.length++
	} else {
		t := l.tail
		newNode := &Node[T]{
			value:    item,
			previous: t,
			next:     nil,
		}
		l.tail.next = newNode
		l.tail = newNode
		l.length++
	}
}

// Adds item to the front of the list
func (l *List[T]) AddFront(item T) {
	if l.head == nil {
		newNode := &Node[T]{
			value:    item,
			previous: nil,
			next:     nil,
		}
		l.head = newNode
		l.tail = newNode
		l.length++
	} else {
		h := l.head
		newNode := &Node[T]{
			value:    item,
			previous: nil,
			next:     h,
		}
		l.head.previous = newNode
		l.head = newNode
		l.length++
	}
}

// Gets item at the front
func (l *List[T]) GetFront() (T, error) {
	var item T
	if l.length == 0 {
		return item, errors.New("empty list")
	}
	return l.head.value, nil
}

// Gets item at the back
func (l *List[T]) GetBack() (T, error) {
	var item T
	if l.length == 0 {
		return item, errors.New("empty list")
	}
	return l.tail.value, nil
}

// Removes head
func (l *List[T]) RemoveFront() (T, error) {
	var item T
	if l.length == 0 {
		return item, errors.New("empty list")
	}
	h := l.head
	item = h.value
	l.head = h.next
	l.head.previous = nil
	return item, nil
}

// Removes tail
func (l *List[T]) RemoveBack() (T, error) {
	var item T
	if l.length == 0 {
		return item, errors.New("empty list")
	}
	t := l.tail
	item = t.value
	l.tail = t.previous
	l.tail.next = nil
	return item, nil
}

// Returns first index of element if exists
func (l *List[T]) IndexOf(value T) int {
	if l.length == 0 {
		return -1
	}
	i := 0
	for n := l.head; n != nil; n = n.next {
		if n.value == value {
			return i
		}
		i++
	}
	return -1
}

// Tests whether the list contains a given value
func (l *List[T]) Contains(value T) bool {
	return l.IndexOf(value) != -1
}

// Returns number of items
func (l *List[T]) Length() int {
	return l.length
}

// Tests whether the lists is empty
func (l *List[T]) IsEmpty() bool {
	return l.length == 0
}

// Empties the list
func (l *List[T]) Clear() {
	l.head = nil
	l.tail = nil
	l.length = 0
}

// Returns a reversed copy of the list
func (l *List[T]) Reverse() *List[T] {
	reverse := NewList[T]()
	for value := range l.Values {
		reverse.AddFront(value)
	}
	return reverse
}

// Returns a copy of the list
func (l *List[T]) Copy() *List[T] {
	copy := NewList[T]()
	for value := range l.Values {
		copy.Add(value)
	}
	return copy
}

// Iterator method
func (l *List[T]) Values(yield func(T) bool) {
	for n := l.head; n != nil; n = n.next {
		if !yield(n.value) {
			return
		}
	}
}

// Reverse iterator method
func (l *List[T]) Backward(yield func(T) bool) {
	for n := l.tail; n != nil; n = n.previous {
		if !yield(n.value) {
			return
		}
	}
}
