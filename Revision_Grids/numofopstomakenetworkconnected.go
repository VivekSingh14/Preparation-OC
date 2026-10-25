package main

import "fmt"

type DFSGraph struct {
	adjlist map[int][]int
}

func NewDFSGraph() *DFSGraph {
	return &DFSGraph{make(map[int][]int)}
}

func (g *DFSGraph) makeAdjList(u int, v int) {
	g.adjlist[u] = append(g.adjlist[u], v)
	g.adjlist[v] = append(g.adjlist[v], u)
}

func main() {
	n := 6
	connections := [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}}

	if n-len(connections) > 1 {
		fmt.Println(-1)
		return
	}

	g := NewDFSGraph()

	for _, val := range connections {
		g.makeAdjList(val[0], val[1])
	}

	visited := make([]int, n)

	components := 0

	for i := 0; i < n; i++ {
		if visited[i] == 0 {
			networkDFS(g, visited, i)

			components++
		}
	}

	fmt.Println(components - 1)
}

func networkDFS(g *DFSGraph, visited []int, source int) {
	visited[source] = 1

	for _, val := range g.adjlist[source] {
		if visited[val] == 0 {
			networkDFS(g, visited, val)
		}
	}
}
