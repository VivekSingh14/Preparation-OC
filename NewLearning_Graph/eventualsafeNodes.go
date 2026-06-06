package main

import "fmt"

func main11() {
	V := 7
	//true case
	edges := [][]int{{0, 1}, {0, 2}, {1, 2}, {1, 3}, {2, 5}, {3, 0}, {4, 5}}

	//false case
	//edges := [][]int{{0, 1}, {0, 2}, {1, 2}, {2, 3}}

	fmt.Println(isCycle3(V, edges))
}

func isCycle3(v int, edges [][]int) []int {
	visited := make([]int, v)
	current_path := make([]int, v)

	adjlist := make([][]int, v)

	for i := 0; i < len(adjlist); i++ {
		adjlist[i] = make([]int, 0)
	}

	for i := 0; i < len(edges); i++ {
		adjlist[edges[i][0]] = append(adjlist[edges[i][0]], edges[i][1])
	}

	for i := 0; i < v; i++ {
		if visited[i] == 0 {
			dfs8(i, adjlist, visited, current_path)

		}
	}
	var safeNodes []int
	for i := 0; i < v; i++ {
		if current_path[i] == 0 {
			safeNodes = append(safeNodes, i)
		}
	}
	return safeNodes
}

func dfs8(node int, edges [][]int, visited []int, current_path []int) bool {

	visited[node] = 1
	current_path[node] = 1

	for _, nbr := range edges[node] {
		if visited[nbr] == 0 {
			ans := dfs8(nbr, edges, visited, current_path)
			if ans {
				return true
			}
		} else {
			if current_path[nbr] == 1 {
				return true
			}
		}
	}
	//very important to reset it again to zero.
	current_path[node] = 0
	return false
}
