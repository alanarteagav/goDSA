package graph

type Graph interface {
	AddEdge(u, v int)
	Colour(v, colour int)
	GetColour(v int) int
	Uncoloured(v int) bool
	Colours() *[]int
	Neighbours(v int) *[]int
}

type AdjacencyMatrixGraph struct {
	adjacencyMatrix [][]int
	colours         []int
}

func NewGraph(n int) Graph {
	matrix := make([][]int, n)
	colours := make([]int, n)
	for i := range n {
		matrix[i] = make([]int, n)
	}
	return &AdjacencyMatrixGraph{
		adjacencyMatrix: matrix,
		colours:         colours,
	}
}

func (g *AdjacencyMatrixGraph) AddEdge(u, v int) {
	g.adjacencyMatrix[u][v] = 1
	g.adjacencyMatrix[v][u] = 1
}

func (g *AdjacencyMatrixGraph) Colour(v, colour int) {
	g.colours[v] = colour
}

func (g *AdjacencyMatrixGraph) GetColour(v int) int {
	return g.colours[v]
}

func (g *AdjacencyMatrixGraph) Uncoloured(v int) bool {
	return g.colours[v] == 0
}

func (g *AdjacencyMatrixGraph) Colours() *[]int {
	return &g.colours
}

func (g *AdjacencyMatrixGraph) Neighbours(v int) *[]int {
	return nil // TODO
}

type AdjacencyListGraph struct {
	adjacencyList [][]int
	colours       []int
}

func NewGraphII(n int) Graph {
	list := make([][]int, n)
	colours := make([]int, n)
	for i := range list {
		list[i] = make([]int, 0)
	}
	return &AdjacencyListGraph{
		adjacencyList: list,
		colours:       colours,
	}
}

func (g *AdjacencyListGraph) AddEdge(u, v int) {
	g.adjacencyList[u] = append(g.adjacencyList[u], v)
	g.adjacencyList[v] = append(g.adjacencyList[v], u)
}

func (g *AdjacencyListGraph) Colour(v, colour int) {
	g.colours[v] = colour
}

func (g *AdjacencyListGraph) GetColour(v int) int {
	return g.colours[v]
}

func (g *AdjacencyListGraph) Uncoloured(v int) bool {
	return g.colours[v] == 0
}

func (g *AdjacencyListGraph) Colours() *[]int {
	return &g.colours
}

func (g *AdjacencyListGraph) Neighbours(v int) *[]int {
	return &g.adjacencyList[v]
}
