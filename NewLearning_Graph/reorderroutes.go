package main

import "fmt"

// https://leetcode.com/problems/reorder-routes-to-make-all-paths-lead-to-the-city-zero/description/
func main7() {
	// n := 6
	// connections := [][]int{{0, 1}, {1, 3}, {2, 3}, {4, 0}, {4, 5}}

	n := 5
	connections := [][]int{{1, 0}, {1, 2}, {3, 2}, {3, 4}}

	fmt.Println(minreOrder(n, connections))

}

func minreOrder(n int, connections [][]int) int {
	visited := make([]int, n)
	forward_nbrs := make([][]int, n)
	backward_nbrs := make([][]int, n)

	for i := 0; i < n; i++ {
		forward_nbrs[i] = make([]int, 5)
		backward_nbrs[i] = make([]int, 5)
	}

	for i := 0; i < len(connections); i++ {
		a := connections[i][0]
		b := connections[i][1]

		forward_nbrs[a] = append(forward_nbrs[a], b)
		backward_nbrs[b] = append(backward_nbrs[b], a)
	}

	ans := 0

	dfs2(0, forward_nbrs, backward_nbrs, &ans, visited)

	return ans

}

func dfs2(source int, forward_nbrs [][]int, backward_nbrs [][]int, ans *int, visited []int) {

	//fmt.Println(source)
	visited[source] = 1

	for i := 0; i < len(forward_nbrs[source]); i++ {
		if visited[forward_nbrs[source][i]] == 0 {
			*ans += 1
			dfs2(forward_nbrs[source][i], forward_nbrs, backward_nbrs, ans, visited)
		}
	}

	for i := 0; i < len(backward_nbrs[source]); i++ {
		if visited[backward_nbrs[source][i]] == 0 {
			dfs2(backward_nbrs[source][i], forward_nbrs, backward_nbrs, ans, visited)
		}
	}

}
