package main

import "fmt"

type cyclicgraph struct {
	adjlist map[int][]int
}

func NewCyclicGraph() *cyclicgraph {
	return &cyclicgraph{make(map[int][]int)}
}

func (g *cyclicgraph) makeAdjList(u int, v int) {
	g.adjlist[u] = append(g.adjlist[u], v)
	g.adjlist[v] = append(g.adjlist[v], u)
}

func main10() {
	n := 6
	edgelist := [][]int{{0, 1}, {1, 2}, {3, 4}, {4, 5}, {5, 3}}

	g := NewCyclicGraph()

	for _, v := range edgelist {
		g.makeAdjList(v[0], v[1])
	}

	fmt.Println(isCycle(n, g.adjlist))

}

func isCycle(n int, adjlist map[int][]int) bool {
	visited := make([]int, n)

	for i := 0; i < n; i++ {
		if visited[i] == 0 {
			ans := cyclicdfs(i, adjlist, visited, -1)
			if ans {
				return true
			}
		}
	}
	return false
}

func cyclicdfs(source int, adjlist map[int][]int, visited []int, parent int) bool {
	visited[source] = 1

	for _, nbr := range adjlist[source] {
		if visited[nbr] == 0 {
			ans := cyclicdfs(nbr, adjlist, visited, source)
			if ans {
				return true
			}
		} else if visited[nbr] == 1 && parent != nbr {
			return true
		}
	}
	return false
}
