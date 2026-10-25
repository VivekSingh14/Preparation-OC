package main

import "fmt"

type UnweightedUnDirectedGraph struct {
	adjlist map[int][]int
}

func UWUDNewGraph() *UnweightedUnDirectedGraph {
	return &UnweightedUnDirectedGraph{adjlist: make(map[int][]int)}
}

func (g *UnweightedUnDirectedGraph) UWUDappend(u int, v int) {
	g.adjlist[u] = append(g.adjlist[u], v)

	//Adding this because its bidirectional
	g.adjlist[v] = append(g.adjlist[v], u)
}

func UWUDtraverse(g *UnweightedUnDirectedGraph) {
	for u, v := range g.adjlist {
		fmt.Println("Vertex: ", u, " is having these neighbours: ", v)
	}
}

func main1() {
	edgelist := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 2}, {1, 3}}

	g1 := UWUDNewGraph()

	for _, v := range edgelist {
		g1.UWUDappend(v[0], v[1])
	}

	UWUDtraverse(g1)

}
