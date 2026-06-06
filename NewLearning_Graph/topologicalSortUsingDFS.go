package main

import "fmt"

var dequeue []int

func main16() {
	//V := 4
	//edges := [][]int{{3, 0}, {1, 0}, {2, 0}}

	V := 6
	edges := [][]int{{1, 3}, {2, 3}, {4, 1}, {4, 0}, {5, 0}, {5, 2}}

	adjlist := make([][]int, V)

	for i := 0; i < len(adjlist); i++ {
		adjlist[i] = make([]int, 0)
	}

	for i := 0; i < len(edges); i++ {
		adjlist[edges[i][0]] = append(adjlist[edges[i][0]], edges[i][1])
	}

	fmt.Println(topologicalSort(adjlist))

}

func topologicalSort(adjlist [][]int) []int {
	n := len(adjlist)

	visited := make([]int, n)
	for i := 0; i < n; i++ {
		if visited[i] == 0 {
			dfs90(i, visited, adjlist)
		}
	}
	var ans []int
	for i := len(dequeue) - 1; i >= 0; i-- {
		ans = append(ans, dequeue[i])
	}
	return ans
}

func dfs90(source int, visited []int, adjlist [][]int) {
	visited[source] = 1
	for _, nbr := range adjlist[source] {
		if visited[nbr] == 0 {
			dfs90(nbr, visited, adjlist)
		}
	}
	dequeue = append(dequeue, source)
}
