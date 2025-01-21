package algorithms

import (
	ds "goDSA/pkg/dataStructures"
	gr "goDSA/pkg/dataStructures/graph"
)

func GraphSearch(G gr.Graph, r int, C ds.Collection[int], colour int) {
	G.Colour(r, colour)
	C.Add(r)
	for !C.IsEmpty() {
		v := C.Remove()
		for _, u := range *G.Neighbours(v) {
			if G.Uncoloured(u) {
				G.Colour(u, colour)
				C.Add(u)
			}
		}
	}
}
