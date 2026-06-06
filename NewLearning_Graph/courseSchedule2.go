package main

import "fmt"

func main8() {
	numCourses := 4
	prerequisites := [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}

	indegrees := indegree2(prerequisites, numCourses)

	adjlist := make([][]int, numCourses)

	for i := 0; i < len(adjlist); i++ {
		adjlist[i] = make([]int, 0)
	}

	for i := 0; i < len(prerequisites); i++ {
		adjlist[prerequisites[i][1]] = append(adjlist[prerequisites[i][1]], prerequisites[i][0])
	}

	fmt.Println(bfs15(indegrees, numCourses, adjlist))
}

func bfs15(indegrees []int, V int, adjlist [][]int) []int {
	var res []int
	var queue []int

	for i := 0; i < len(indegrees); i++ {
		if indegrees[i] == 0 {
			queue = append(queue, i)
		}
	}

	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
		res = append(res, front)
		for _, nbr := range adjlist[front] {
			indegrees[nbr]--

			if indegrees[nbr] == 0 {
				queue = append(queue, nbr)
			}
		}
	}
	return res
}

func indegree2(edges [][]int, v int) []int {
	indegrees := make([]int, v)
	for i := 0; i < len(edges); i++ {
		indegrees[edges[i][0]] += 1
	}
	return indegrees
}
