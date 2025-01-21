package algorithms

import (
	ds "github.com/alanarteagav/goDSA/pkg/dataStructures"
	gr "github.com/alanarteagav/goDSA/pkg/dataStructures/graph"
)

func DFS(G gr.Graph, r int) {
	S := new(ds.Stack[int])
	G.Colour(r, 1)
	S.Push(r)
	for !S.IsEmpty() {
		v := S.Pop()
		for _, u := range *G.Neighbours(v) {
			if G.Uncoloured(u) {
				G.Colour(u, 1)
				S.Push(u)
			}
		}
	}
}
