package algorithms

import (
	ds "github.com/alanarteagav/goDSA/pkg/dataStructures"
	gr "github.com/alanarteagav/goDSA/pkg/dataStructures/graph"
)

func CountConnectedComponents(G gr.Graph) int {
	colours := G.Colours()
	totalColours := 0
	for i, colour := range *colours {
		if colour == 0 {
			GraphSearch(G, i, new(ds.Stack[int]), totalColours+1)
			totalColours++
		}
	}
	return totalColours
}
