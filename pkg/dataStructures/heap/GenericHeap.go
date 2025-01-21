package heap

import ds "goDSA/pkg/dataStructures"

type HeapItem[V any, P ds.Ordered] struct {
	Value    V // The value of the item; arbitrary.
	Priority P // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

func NewHeapItem[V any, P ds.Ordered](value V, priority P) HeapItem[V, P] {
	return HeapItem[V, P]{
		Value:    value,
		Priority: priority,
	}
}

type GenericHeap[T1 any, T2 ds.Ordered] []*HeapItem[T1, T2]

func (h GenericHeap[T1, T2]) Len() int {
	return len(h)
}

func (h GenericHeap[T1, T2]) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}

func (h GenericHeap[T1, T2]) Less(i, j int) bool {
	return h[i].Priority < h[j].Priority
}

func (h *GenericHeap[T1, T2]) Push(x any) {
	n := len(*h)
	item := x.(*HeapItem[T1, T2])
	item.index = n
	*h = append(*h, item)
}

func (h *GenericHeap[T1, T2]) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1 // for safety
	*h = old[0 : n-1]
	return item
}
