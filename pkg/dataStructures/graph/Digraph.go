package graph

type Digraph interface {
	AddEdge(u, v int)
	Colour(v, colour int)
	GetColour(v int) int
	Uncoloured(v int) bool
	Colours() *[]int
	Neighbours(v int) *[]int
}
