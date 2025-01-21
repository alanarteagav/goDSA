package algorithms

import (
	ds "github.com/alanarteagav/goDSA/pkg/dataStructures"
	gr "github.com/alanarteagav/goDSA/pkg/dataStructures/graph"
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
