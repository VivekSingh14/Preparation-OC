package main

import "fmt"

// Unweighted Unidirectional graph
type CSBFSGraph struct {
	adjlist map[int][]int
}

func NewCSBFSGraph() *CSBFSGraph {
	return &CSBFSGraph{adjlist: make(map[int][]int)}
}

func (g *CSBFSGraph) CSBFSAppend(u int, v int) {

	//interchanges because v is dependent on u, as given in question
	g.adjlist[v] = append(g.adjlist[v], u)
}

func CSTopologicalSortBFS(s *CSBFSGraph, n int) []int {

	indegree := make([]int, n)
	res := make([]int, 0)

	for _, neighbors := range s.adjlist {

		for _, v := range neighbors {
			indegree[v]++
		}
	}
	var que []int
	for i := 0; i < len(indegree); i++ {
		if indegree[i] == 0 {
			que = append(que, i)
		}
	}

	for len(que) != 0 {
		node := que[0]
		que = que[1:]

		res = append(res, node)

		for _, nbr := range s.adjlist[node] {
			indegree[nbr]--
			if indegree[nbr] == 0 {
				que = append(que, nbr)
			}
		}
	}

	//extra loop to check if cycle exist ?
	for i := 0; i < len(indegree); i++ {
		if indegree[i] != 0 {
			return make([]int, 0)
		}
	}

	return res

}

func main20() {

	n := 4

	prerequisites := [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}

	g1 := NewCSBFSGraph()

	for _, v := range prerequisites {
		g1.CSBFSAppend(v[0], v[1])
	}

	fmt.Println(CSTopologicalSortBFS(g1, n))

}
