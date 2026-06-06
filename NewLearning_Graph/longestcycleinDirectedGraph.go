package main

import "fmt"

func main12() {

	//edges := []int{3, 3, 4, 2, 3}
	edges := []int{2, -1, 3, 1}

	V := len(edges)

	fmt.Println(longestCycle(V, edges))

}

func longestCycle(v int, edges []int) int {
	visited := make([]int, v)
	current_path := make([]int, v)
	longest_path := -1
	adjlist := make([][]int, v)
	for i := 0; i < len(adjlist); i++ {
		adjlist[i] = make([]int, 0)
	}
	for i := 0; i < len(edges); i++ {
		if edges[i] >= 0 {
			adjlist[i] = append(adjlist[i], edges[i])
		}
	}

	for i := 0; i < v; i++ {
		if visited[i] == 0 {
			dfs9(i, adjlist, visited, current_path, &longest_path, 1)

		}
	}
	return longest_path
}

func dfs9(node int, edges [][]int, visited []int, current_path []int, longest_path *int, path_pos int) bool {

	visited[node] = 1
	current_path[node] = path_pos

	for _, nbr := range edges[node] {
		if visited[nbr] == 0 {
			ans := dfs9(nbr, edges, visited, current_path, longest_path, path_pos+1)
			if ans {
				return true
			}
		} else if current_path[nbr] > 1 {
			cycle_length := path_pos - current_path[nbr] + 1
			if cycle_length > *longest_path {
				*longest_path = cycle_length
			}
			return true
		}
	}

	//very important to reset it again to zero.
	current_path[node] = 0
	return false
}
