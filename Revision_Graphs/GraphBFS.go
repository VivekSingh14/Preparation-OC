package main

import "fmt"

// Unweighted Bidirectional graph
type BFSGraph struct {
	adjlist map[int][]int
}

func NewBFSGraph() *BFSGraph {
	return &BFSGraph{adjlist: make(map[int][]int)}
}

func (g *BFSGraph) BFSAppend(u int, v int) {
	g.adjlist[u] = append(g.adjlist[u], v)
	g.adjlist[v] = append(g.adjlist[v], u)
}

func displayNeighbours(g *BFSGraph) {
	for u, v := range g.adjlist {
		fmt.Println("Vertex is: ", u, "and neighbours are: ", v)
	}
}

func BFS(s *BFSGraph, source int) {
	var que []int
	que = append(que, source)
	visited := make([]int, 7)

	for len(que) != 0 {
		node := que[0]
		que = que[1:]

		fmt.Print(node, " ")

		visited[node] = 1

		if val, ok := s.adjlist[node]; ok {
			for _, neighbour := range val {
				if visited[neighbour] == 0 {
					que = append(que, neighbour)
					visited[neighbour] = 1
				}
			}
		}
	}

	fmt.Println()

}

func main4() {
	edgelist := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 2}, {1, 3}, {2, 6}, {3, 6}, {6, 5}}
	g1 := NewBFSGraph()

	for _, v := range edgelist {
		g1.BFSAppend(v[0], v[1])
	}

	displayNeighbours(g1)

	fmt.Println("*****************")

	BFS(g1, 1)

}
