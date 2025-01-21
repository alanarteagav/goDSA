package graph

type WeightedEdge struct {
	Vertex int
	Weight float64
}

type WeightedDigraph interface {
	AddArc(u, v int, w float64)
	Colour(v, colour int)
	GetColour(v int) int
	Uncoloured(v int) bool
	Colours() *[]int
	Neighbours(v int) *[]WeightedEdge
	Order() int
}

type AdjacencyListDigraph struct {
	adjacencyList [][]WeightedEdge
	colours       []int
}

func NewAdjacencyListDigraph(n int) WeightedDigraph {
	list := make([][]WeightedEdge, n)
	colours := make([]int, n)
	for i := range list {
		list[i] = make([]WeightedEdge, 0)
	}
	return &AdjacencyListDigraph{
		adjacencyList: list,
		colours:       colours,
	}
}

func (d *AdjacencyListDigraph) AddArc(u, v int, w float64) {
	d.adjacencyList[u] = append(d.adjacencyList[u], WeightedEdge{
		Vertex: v,
		Weight: w,
	})
}

func (d *AdjacencyListDigraph) Colour(v, colour int) {
	d.colours[v] = colour
}

func (d *AdjacencyListDigraph) GetColour(v int) int {
	return d.colours[v]
}

func (d *AdjacencyListDigraph) Uncoloured(v int) bool {
	return d.colours[v] == 0
}

func (d *AdjacencyListDigraph) Colours() *[]int {
	return &d.colours
}

func (d *AdjacencyListDigraph) Neighbours(v int) *[]WeightedEdge {
	return &d.adjacencyList[v]
}

func (d *AdjacencyListDigraph) Order() int {
	return len(d.adjacencyList)
}

func (d *AdjacencyListDigraph) AdjacencyList() [][]WeightedEdge {
	return d.adjacencyList
}
