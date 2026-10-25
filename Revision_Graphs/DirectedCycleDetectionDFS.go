package main

import "fmt"

type cyclicdirgraph struct {
	adjlist map[int][]int
}

func NewCyclicDIRGraph() *cyclicdirgraph {
	return &cyclicdirgraph{make(map[int][]int)}
}

func (g *cyclicdirgraph) makeAdjList(u int, v int) {
	g.adjlist[u] = append(g.adjlist[u], v)
	//g.adjlist[v] = append(g.adjlist[v], u)
}

func main12() {
	n := 8
	edgelist := [][]int{{0, 1}, {1, 2}, {3, 4}, {4, 5}, {5, 3}, {6, 7}, {2, 0}}

	g := NewCyclicDIRGraph()

	for _, v := range edgelist {
		g.makeAdjList(v[0], v[1])
	}

	fmt.Println(isCycleDIR(n, g.adjlist))

}

func isCycleDIR(n int, adjlist map[int][]int) bool {
	visited := make([]int, n)
	current_path := make([]int, n)

	for i := 0; i < n; i++ {
		if visited[i] == 0 {
			ans := cyclicDIRdfs(i, adjlist, visited, current_path)
			if ans {
				return true
			}
		}
	}
	return false
}

func cyclicDIRdfs(source int, adjlist map[int][]int, visited []int, current_path []int) bool {
	visited[source] = 1
	current_path[source] = 1
	for _, nbr := range adjlist[source] {
		if visited[nbr] == 0 {
			ans := cyclicDIRdfs(nbr, adjlist, visited, current_path)
			if ans {
				return true
			}
		} else if visited[nbr] == 1 && current_path[nbr] == 1 {
			return true
		}
	}
	current_path[source] = 0
	return false
}
