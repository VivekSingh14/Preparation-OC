package main

import "fmt"

// Unweighted Bidirectional graph
type DFSGraph struct {
	adjlist map[int][]int
}

func NewDFSGraph() *DFSGraph {
	return &DFSGraph{adjlist: make(map[int][]int)}
}

func (g *DFSGraph) DFSAppend(u int, v int) {
	g.adjlist[u] = append(g.adjlist[u], v)
	g.adjlist[v] = append(g.adjlist[v], u)
}

func displayNeighbour(g *DFSGraph) {
	for u, v := range g.adjlist {
		fmt.Println("Vertex is: ", u, "and neighbours are: ", v)
	}
}

func DFS(s *DFSGraph, source int, visited []int) {
	fmt.Println(source)
	visited[source] = 1

	if val, ok := s.adjlist[source]; ok {
		for _, neighbour := range val {
			if visited[neighbour] == 0 {
				DFS(s, neighbour, visited)
			}
		}
	}
}

func main5() {
	edgelist := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 2}, {1, 3}, {2, 6}, {3, 6}, {6, 5}}
	g1 := NewDFSGraph()

	for _, v := range edgelist {
		g1.DFSAppend(v[0], v[1])
	}

	displayNeighbour(g1)

	fmt.Println("*****************")

	visited := make([]int, 7)
	DFS(g1, 1, visited)

}
