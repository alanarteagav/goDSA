package main

import (
	"fmt"
	"math/rand/v2"

	algo "github.com/alanarteagav/goDSA/pkg/algorithms"
	graphAlgo "github.com/alanarteagav/goDSA/pkg/algorithms/graph"
	binaryTree "github.com/alanarteagav/goDSA/pkg/dataStructures/binaryTree"
	graph "github.com/alanarteagav/goDSA/pkg/dataStructures/graph"
	hp "github.com/alanarteagav/goDSA/pkg/dataStructures/heap"
	list "github.com/alanarteagav/goDSA/pkg/dataStructures/list"
)

func testList() {
	l := list.NewList[int]()
	l.Add(0)
	l.Add(1)
	l.AddFront(2)
	r := l.Reverse()
	c := r.Copy()
	for v := range c.Values {
		fmt.Println(v)
	}
	l.Clear()
	fmt.Println(l.IsEmpty())
}

func testBinarySearch() {
	intSlice := make([]int, 19)
	for i := range intSlice {
		intSlice[i] = rand.IntN(100)
	}
	intSlice = append(intSlice, 42)
	algo.QuickSort(&intSlice, 0, len(intSlice)-1)
	result := algo.BinarySearch(intSlice, 42)
	fmt.Println(intSlice)
	fmt.Println(result)
}

func testMergeSort() {
	intSlice := make([]int, 19)
	for i := range intSlice {
		intSlice[i] = rand.IntN(100)
	}
	fmt.Println(intSlice)
	algo.MergeSort(&intSlice, 0, len(intSlice)-1)
	fmt.Println(intSlice)
}

func testQuickSort() {
	intSlice := make([]int, 19)
	for i := range intSlice {
		intSlice[i] = rand.IntN(100)
	}
	fmt.Println(intSlice)
	algo.QuickSort(&intSlice, 0, len(intSlice)-1)
	fmt.Println(intSlice)
}

func testBinaryTree() {
	t := binaryTree.NewBinarySearchTree[int]()
	// t = new(binaryTree.BinarySearchTree[int])
	t.AddIterative(2)
	t.AddIterative(3)
	t.AddIterative(1)
	fmt.Println(t.PreorderTraversal())
	fmt.Println(t.InorderTraversal())
	fmt.Println(t.PostorderTraversal())
	t.Delete(1)
	fmt.Println(t.InorderTraversal())
	t.Delete(2)
	fmt.Println(t.InorderTraversal())
}

func testGraph() {
	graph := graph.NewGraphII(12)
	graph.AddEdge(0, 11)
	graph.AddEdge(2, 11)
	graph.AddEdge(3, 11)
	graph.AddEdge(1, 10)
	//graph.AddEdge(10, 11)
	// fmt.Println(graph)
	// graphAlgo.BFS(graph, 0)
	// graphAlgo.DFS(graph, 0)
	//graphAlgo.GraphSearch(graph, 0, new(ds.Stack[int]), 2)
	fmt.Println(graphAlgo.CountConnectedComponents(graph))
	fmt.Println(graph.Colours())
}

func testHeap() {
	slc := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
	fmt.Println(slc)

	hp := hp.NewMinHeap[int]()
	hp.BuildMaxHeap(&slc)
	/*
		fmt.Println(slc)

		fmt.Println(hp.Pop())
		fmt.Println(slc)
		fmt.Println(hp.Length())
	*/
	fmt.Println(hp.HeapSort())
}

func testGenericHeap() {
	/*
		heap := make(hp.GenericHeap[int, int], 0)
		letters := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
		for _, v := range letters {
			heap = append(heap, hp.HeapItem[int, int]{
				Value:    v,
				Priority: v,
			})
		}
		pq := hp.NewPriorityQueue[int, int]()
		pq.BuildHeap(heap)
		fmt.Println(heap)
	*/

	numbers := []int{2, 1, 5, 3}
	pqII := hp.NewPriorityQueue[int, int]()
	for _, v := range numbers {
		pqII.Push(&hp.HeapItem[int, int]{
			Value:    v,
			Priority: v,
		})
	}
	fmt.Println(pqII.Heap())
	for pqII.Length() > 0 {
		fmt.Println(pqII.Pop())
	}
}

func testDijkstra() {
	weightedDigraph := graph.NewAdjacencyListDigraph(5)
	weightedDigraph.AddArc(0, 1, 10)
	weightedDigraph.AddArc(0, 3, 5)
	weightedDigraph.AddArc(1, 2, 1)
	weightedDigraph.AddArc(1, 3, 2)
	weightedDigraph.AddArc(2, 4, 4)
	weightedDigraph.AddArc(3, 1, 3)
	weightedDigraph.AddArc(3, 2, 9)
	weightedDigraph.AddArc(3, 4, 2)
	weightedDigraph.AddArc(4, 1, 7)
	weightedDigraph.AddArc(4, 2, 6)

	fmt.Println(graphAlgo.Dijkstra(weightedDigraph, 0))
}

func main() {
	// testList()
	//testBinaryTree()
	// testGraph()

	// testBinarySearch()
	// testQuickSort()
	// testMergeSort()

	// testHeap()
	// testGenericHeap()
	testDijkstra()
}
