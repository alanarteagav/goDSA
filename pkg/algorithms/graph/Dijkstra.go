package algorithms

import (
	gr "goDSA/pkg/dataStructures/graph"
	heap "goDSA/pkg/dataStructures/heap"
	"math"
)

func Dijkstra(D gr.WeightedDigraph, r int) ([]int, []float64) {
	pointers := make([]*heap.HeapItem[int, float64], D.Order())
	PQ := heap.NewPriorityQueue[int, float64]()

	distance := make([]float64, D.Order())
	parent := make([]int, D.Order())

	rootItem := heap.NewHeapItem(r, 0.0)
	pointers[r] = &rootItem
	PQ.Push(&rootItem)
	parent[0] = -1

	for v := range D.Order() {
		if v != r {
			item := heap.NewHeapItem(v, math.Inf(1))
			pointers[v] = &item
			PQ.Push(&item)
			parent[v] = -1
		}
	}

	for !PQ.IsEmpty() {
		pair, _ := PQ.Pop()
		v := pair.Value
		D.Colour(v, 1)
		for _, e := range *D.Neighbours(v) {
			u := e.Vertex
			vDistance := (*pointers[v]).Priority
			uDistance := (*pointers[u]).Priority
			newDistance := vDistance + e.Weight
			if D.Uncoloured(u) && uDistance > newDistance {
				parent[u] = v
				PQ.Update(pointers[u], pointers[u].Value, newDistance)
				distance[u] = newDistance
			}
		}
	}
	return parent, distance
}
