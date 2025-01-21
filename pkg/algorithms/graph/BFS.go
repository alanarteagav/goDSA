package algorithms

import (
	ds "goDSA/pkg/dataStructures"
	gr "goDSA/pkg/dataStructures/graph"
)

func BFS(G gr.Graph, r int) {
	Q := new(ds.Queue[int])
	G.Colour(r, 1)
	Q.Enqueue(r)
	for !Q.IsEmpty() {
		v := Q.Dequeue()
		for _, u := range *G.Neighbours(v) {
			if G.Uncoloured(u) {
				G.Colour(u, 1)
				Q.Enqueue(u)
			}
		}
	}
}
