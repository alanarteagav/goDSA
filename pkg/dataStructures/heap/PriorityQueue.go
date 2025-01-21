package heap

import (
	"container/heap"
	hp "container/heap"
	"errors"
	ds "goDSA/pkg/dataStructures"
)

type PriorityQueue[V any, P ds.Ordered] struct {
	heap GenericHeap[V, P]
}

func NewPriorityQueue[V any, P ds.Ordered]() *PriorityQueue[V, P] {
	heap := make([]*HeapItem[V, P], 0)
	return &PriorityQueue[V, P]{
		heap: heap,
	}
}

func (pq *PriorityQueue[V, P]) BuildHeap(slc GenericHeap[V, P]) {
	pq.heap = slc
	hp.Init(&pq.heap)
}

func (pq *PriorityQueue[V, P]) Heap() *GenericHeap[V, P] {
	return &pq.heap
}

func (pq *PriorityQueue[V, P]) Push(item *HeapItem[V, P]) {
	hp.Push(&pq.heap, item)
}

func (pq *PriorityQueue[V, P]) Length() int {
	return len(pq.heap)
}

func (pq *PriorityQueue[V, P]) IsEmpty() bool {
	return len(pq.heap) == 0
}

func (pq *PriorityQueue[V, P]) Peek() (*HeapItem[V, P], error) {
	var item *HeapItem[V, P]
	if len(pq.heap) != 0 {
		return pq.heap[0], nil
	}
	return item, errors.New("empy priority queue")
}

func (pq *PriorityQueue[V, P]) Pop() (*HeapItem[V, P], error) {
	var item *HeapItem[V, P]
	if len(pq.heap) != 0 {
		return hp.Pop(&pq.heap).(*HeapItem[V, P]), nil
	}
	return item, errors.New("empy priority queue")
}

func (pq *PriorityQueue[V, P]) Update(item *HeapItem[V, P], value V, priority P) {
	item.Value = value
	item.Priority = priority
	heap.Fix(&pq.heap, item.index)
}
