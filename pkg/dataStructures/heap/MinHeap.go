package heap

import (
	"errors"
	ds "goDSA/pkg/dataStructures"
)

type MinHeap[T ds.Ordered] struct {
	heap *[]T
}

func NewMinHeap[T ds.Ordered]() Heap[T] {
	heap := make([]T, 0)
	return &MinHeap[T]{
		heap: &heap,
	}
}

func (h *MinHeap[T]) left(i int) int {
	return 2*i + 1
}

func (h *MinHeap[T]) right(i int) int {
	return 2*i + 2
}

func (h *MinHeap[T]) heapify(i int) {
	length := len(*h.heap)
	l := h.left(i)
	r := h.right(i)
	smallest := i
	if l < length && (*h.heap)[l] < (*h.heap)[i] {
		smallest = l
	}
	if r < length && (*h.heap)[r] < (*h.heap)[smallest] {
		smallest = r
	}
	if smallest != i {
		(*h.heap)[i], (*h.heap)[smallest] = (*h.heap)[smallest], (*h.heap)[i]
		h.heapify(smallest)
	}
}

func (h *MinHeap[T]) BuildMaxHeap(slc *[]T) {
	h.heap = slc
	for i := len(*h.heap)/2 - 1; i >= 0; i-- {
		h.heapify(i)
	}
}

// TODO : improve inplace
func (h *MinHeap[T]) HeapSort() []T {
	sorted := make([]T, len(*h.heap))
	n := len(*h.heap)
	for i := 0; i < n; i++ {
		sorted[i], _ = h.Pop()
	}
	return sorted
}

func (h *MinHeap[T]) Insert(item T) {}

func (h *MinHeap[T]) Length() int {
	return len(*h.heap)

}
func (h *MinHeap[T]) Peek() (T, error) {
	var item T
	if len(*h.heap) > 0 {
		return (*h.heap)[0], nil
	}
	return item, errors.New("empty heap")
}

func (h *MinHeap[T]) Pop() (T, error) {
	var item T
	if len(*h.heap) > 0 {
		min := (*h.heap)[0]
		(*h.heap)[0] = (*h.heap)[len(*h.heap)-1]
		*h.heap = (*h.heap)[0 : len(*h.heap)-1]
		h.heapify(0)
		return min, nil
	}
	return item, errors.New("empty heap")
}
