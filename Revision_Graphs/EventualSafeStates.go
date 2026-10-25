package main

import "fmt"

type cyclicdirSSgraph struct {
	adjlist map[int][]int
}

func NewCyclicDIRSSGraph() *cyclicdirSSgraph {
	return &cyclicdirSSgraph{make(map[int][]int)}
}

func (g *cyclicdirSSgraph) makeAdjList(u int, v int) {
	g.adjlist[u] = append(g.adjlist[u], v)
}

func main14() {
	n := 6
	edgelist := [][]int{{0, 1}, {1, 2}, {2, 0}, {0, 3}, {3, 4}, {5, 4}}

	g := NewCyclicDIRSSGraph()

	for _, v := range edgelist {
		g.makeAdjList(v[0], v[1])
	}

	res := isCycleSSDIR(n, g.adjlist)

	for val := range res {
		if res[val] == 0 {
			fmt.Print(val, " ")
		}
	}
}

func isCycleSSDIR(n int, adjlist map[int][]int) []int {
	visited := make([]int, n)
	current_path := make([]int, n)

	for i := 0; i < n; i++ {
		if visited[i] == 0 {
			ans := cyclicDIRSSdfs(i, adjlist, visited, current_path)
			if ans {
				return current_path
			}
		}
	}
	return nil
}

func cyclicDIRSSdfs(source int, adjlist map[int][]int, visited []int, current_path []int) bool {
	visited[source] = 1
	current_path[source] = 1
	for _, nbr := range adjlist[source] {
		if visited[nbr] == 0 {
			ans := cyclicDIRSSdfs(nbr, adjlist, visited, current_path)
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
