package main

import "fmt"

type UnWeightedUniDirectionGraph struct {
	adjlist map[int][]int
}

func UWUNIDNewGraph() *UnWeightedUniDirectionGraph {
	return &UnWeightedUniDirectionGraph{adjlist: make(map[int][]int)}
}

func (g *UnWeightedUniDirectionGraph) uwuniappend(u int, v int) {

	// only once because it is uni directional graph
	g.adjlist[u] = append(g.adjlist[u], v)
}

func UWUNItraverse(g *UnWeightedUniDirectionGraph) {
	for u, v := range g.adjlist {
		fmt.Println("Vertex: ", u, "is having ", v, "neighbours")
	}
}

func main2() {
	edgelist := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 2}, {1, 3}}

	g2 := UWUNIDNewGraph()

	for _, v := range edgelist {
		g2.uwuniappend(v[0], v[1])
	}

	UWUNItraverse(g2)

}

/*------
-------right now output is---------

------Where you can see vertex are not in sequence---------

Vertex:  3 is having  [4] neighbours
Vertex:  4 is having  [2] neighbours
Vertex:  1 is having  [2 3] neighbours
Vertex:  2 is having  [3] neighbours


-------*/
