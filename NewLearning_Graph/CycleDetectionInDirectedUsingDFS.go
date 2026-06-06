package main

import "fmt"

func main10() {
	V := 4
	//true case
	edges := [][]int{{0, 1}, {0, 2}, {1, 2}, {2, 3}, {2, 0}}

	//false case
	//edges := [][]int{{0, 1}, {0, 2}, {1, 2}, {2, 3}}

	fmt.Println(isCycle2(V, edges))
}

func isCycle2(v int, edges [][]int) bool {
	visited := make([]int, v)
	current_path := make([]int, v)
	// for i := 0; i < len(current_path); i++ {
	// 	current_path[i] = -1
	// }

	adjlist := make([][]int, v)

	for i := 0; i < len(adjlist); i++ {
		adjlist[i] = make([]int, 0)
	}

	for i := 0; i < len(edges); i++ {
		adjlist[edges[i][0]] = append(adjlist[edges[i][0]], edges[i][1])
		//adjlist[edges[i][1]] = append(adjlist[edges[i][1]], edges[i][0])
	}

	for i := 0; i < v; i++ {
		if visited[i] == 0 {
			ans := dfs7(i, adjlist, visited, current_path)
			if ans {
				return true
			}

		}
	}
	return false
}

func dfs7(node int, edges [][]int, visited []int, current_path []int) bool {

	visited[node] = 1
	current_path[node] = 1

	for _, nbr := range edges[node] {
		if visited[nbr] == 0 {
			ans := dfs7(nbr, edges, visited, current_path)
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
