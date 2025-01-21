package heap

import ds "goDSA/pkg/dataStructures"

type Heap[T ds.Ordered] interface {
	BuildMaxHeap(*[]T)
	HeapSort() []T
	Insert(item T)
	Length() int
	Peek() (T, error)
	Pop() (T, error)
}
