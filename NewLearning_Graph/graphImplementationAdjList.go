package main

import "fmt"

// edge list is given:
// {1,2}, {2,3}, {3,4}, {4,2}, {1,3}
type Graph struct {
	adjlist map[int][]int
}

func NewGraph() *Graph {
	return &Graph{adjlist: make(map[int][]int)}
}

func (g *Graph) addEdges(u int, v int) {
	g.adjlist[u] = append(g.adjlist[u], v)
	//if graph is undirected
	g.adjlist[v] = append(g.adjlist[v], u)
}

func traverseGraph(g *Graph) {

	for u, v := range g.adjlist {

		fmt.Println("vertex: ", u, "and its neighbours are: ", v)
	}

}

func main1() {

	g := NewGraph()

	g.addEdges(1, 2)
	g.addEdges(2, 3)
	g.addEdges(3, 4)
	g.addEdges(4, 2)
	g.addEdges(1, 3)

	traverseGraph(g)

}
