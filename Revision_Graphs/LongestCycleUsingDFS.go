package main

import "fmt"

var longestCycleLen = -1

type longestcyclicdirgraph struct {
	adjlist map[int][]int
}

func NewLongestCyclicDIRGraph() *longestcyclicdirgraph {
	return &longestcyclicdirgraph{make(map[int][]int)}
}

func (g *longestcyclicdirgraph) makeLongestAdjList(u int, v int) {
	g.adjlist[u] = append(g.adjlist[u], v)
}

func main15() {
	//n := 8
	edgelist := [][]int{{0, 1}, {1, 2}, {1, 3}, {3, 4}, {2, 5}, {5, 6}, {6, 7}, {7, 1}}

	g := NewLongestCyclicDIRGraph()

	for _, v := range edgelist {
		g.makeLongestAdjList(v[0], v[1])
	}

	fmt.Println(longestCycle(g.adjlist))
}

func longestCycle(adjlist map[int][]int) int {

	n := len(adjlist)
	visited := make([]int, n+1)
	current_path := make([]int, n+1)

	for i := 0; i < n; i++ {
		if visited[i] == 0 {
			dfsFindLongestCycle(0, i, adjlist, visited, current_path)

		}
	}
	return longestCycleLen
}

func dfsFindLongestCycle(cyclelen int, source int, adjlist map[int][]int, visited []int, current_path []int) {
	cyclelen++
	current_path[source] = cyclelen
	visited[source] = 1

	for _, nbr := range adjlist[source] {
		if visited[nbr] == 0 {
			dfsFindLongestCycle(cyclelen, nbr, adjlist, visited, current_path)

		} else if visited[nbr] == 1 && current_path[nbr] != 0 {
			current_cyclelen := current_path[source] - current_path[nbr] + 1
			longestCycleLen = max(longestCycleLen, current_cyclelen)
		}
	}
	// Remove node from current DFS path.
	current_path[source] = 0
}
