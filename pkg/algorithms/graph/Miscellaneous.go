package algorithms

import (
	ds "goDSA/pkg/dataStructures"
	gr "goDSA/pkg/dataStructures/graph"
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
