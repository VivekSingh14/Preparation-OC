package main

import "fmt"

// Unweighted Unidirectional graph
type TopoDFSGraph struct {
	adjlist map[int][]int
}

func NewTopoDFSGraph() *TopoDFSGraph {
	return &TopoDFSGraph{adjlist: make(map[int][]int)}
}

func (g *TopoDFSGraph) TopoDFSAppend(u int, v int) {
	g.adjlist[u] = append(g.adjlist[u], v)
}

var que []int

func topologicSort(g *TopoDFSGraph, n int) []int {
	visited := make([]int, n)

	for i := 1; i < n; i++ {
		if visited[i] == 0 {
			topodfs(i, visited, g.adjlist)
		}
	}

	return que

}

func topodfs(source int, visited []int, adjlist map[int][]int) {

	visited[source] = 1

	if val, ok := adjlist[source]; ok {
		for _, neighbour := range val {
			if visited[neighbour] == 0 {
				topodfs(neighbour, visited, adjlist)
			}
		}
	}

	que = append(que, source)
}

func main19() {
	n := 8
	edgelist := [][]int{{1, 3}, {2, 3}, {3, 4}, {3, 5}, {4, 7}, {5, 7}, {6, 7}}
	g1 := NewTopoDFSGraph()

	for _, v := range edgelist {
		g1.TopoDFSAppend(v[0], v[1])
	}

	res := topologicSort(g1, n)

	for i := len(res) - 1; i >= 0; i-- {
		fmt.Print(res[i], " ")
	}
}
