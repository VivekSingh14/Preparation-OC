package main

import "fmt"

func main15() {

	V := 4
	//false case
	//edges := [][]int{{3, 0}, {1, 0}, {2, 0}}

	// true case
	edges := [][]int{{0, 1}, {0, 2}, {1, 2}, {2, 3}, {2, 0}}

	indegrees := indegree1(edges, V)

	adjlist := make([][]int, V)

	for i := 0; i < len(adjlist); i++ {
		adjlist[i] = make([]int, 0)
	}

	for i := 0; i < len(edges); i++ {
		adjlist[edges[i][0]] = append(adjlist[edges[i][0]], edges[i][1])
	}

	fmt.Println("is cyclic: ", bfs14(indegrees, V, adjlist))

}

func bfs14(indegrees []int, v int, adjlist [][]int) bool {

	var queue []int
	for i := 0; i < len(indegrees); i++ {
		if indegrees[i] == 0 {
			queue = append(queue, i)
		}
	}

	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
		for _, nbr := range adjlist[front] {
			indegrees[nbr]--
			if indegrees[nbr] == 0 {
				queue = append(queue, nbr)
			}
		}
	}

	for i := 0; i < len(indegrees); i++ {
		if indegrees[i] > 0 {
			return true
		}
	}
	return false
}

func indegree1(edges [][]int, v int) []int {
	indegrees := make([]int, v)
	for i := 0; i < len(edges); i++ {
		indegrees[edges[i][1]] += 1
	}
	return indegrees
}
