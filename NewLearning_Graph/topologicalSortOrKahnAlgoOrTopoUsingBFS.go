package main

import "fmt"

func main14() {

	V := 4
	edges := [][]int{{3, 0}, {1, 0}, {2, 0}}

	indegrees := indegree(edges, V)

	adjlist := make([][]int, V)

	for i := 0; i < len(adjlist); i++ {
		adjlist[i] = make([]int, 0)
	}

	for i := 0; i < len(edges); i++ {
		adjlist[edges[i][0]] = append(adjlist[edges[i][0]], edges[i][1])
	}

	fmt.Println(bfs12(indegrees, V, adjlist))

}

func bfs12(indegrees []int, v int, adjlist [][]int) []int {
	var ans []int

	var queue []int
	for i := 0; i < len(indegrees); i++ {
		if indegrees[i] == 0 {
			queue = append(queue, i)
		}
	}

	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
		ans = append(ans, front)
		for _, nbr := range adjlist[front] {
			indegrees[nbr]--
			if indegrees[nbr] == 0 {
				queue = append(queue, nbr)
			}
		}
	}

	return ans
}

func indegree(edges [][]int, v int) []int {
	indegrees := make([]int, v)
	for i := 0; i < len(edges); i++ {
		indegrees[edges[i][1]] += 1
	}
	return indegrees
}
