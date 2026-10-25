package main

import "fmt"

type cyclicBFSgraph struct {
	adjlist map[int][]int
}

func NewCyclicBFSGraph() *cyclicBFSgraph {
	return &cyclicBFSgraph{make(map[int][]int)}
}

func (g *cyclicBFSgraph) makeAdjList(u int, v int) {
	g.adjlist[u] = append(g.adjlist[u], v)
	g.adjlist[v] = append(g.adjlist[v], u)
}

func main11() {
	n := 6
	edgelist := [][]int{{0, 1}, {1, 2}, {3, 4}, {4, 5}, {5, 3}}

	g := NewCyclicBFSGraph()

	for _, v := range edgelist {
		g.makeAdjList(v[0], v[1])
	}

	fmt.Println(isCycleBFS(n, g.adjlist))

}

func isCycleBFS(n int, adjlist map[int][]int) bool {
	visited := make([]int, n)

	parent := make([]int, n)

	for i := 0; i < n; i++ {
		parent[i] = -1
	}

	for i := 0; i < n; i++ {
		if visited[i] == 0 {
			ans := cyclicbfs(i, adjlist, visited, parent)
			if ans {
				return true
			}
		}
	}
	return false
}

func cyclicbfs(source int, adjlist map[int][]int, visited []int, parent []int) bool {

	var que []int

	que = append(que, source)
	visited[source] = 1

	for len(que) > 0 {
		node := que[0]
		que = que[1:]

		for _, nbr := range adjlist[node] {
			if visited[nbr] == 0 {
				visited[nbr] = 1
				parent[nbr] = node
				que = append(que, nbr)
			} else {
				if parent[node] != nbr {
					return true
				}
			}
		}
	}

	return false
}
