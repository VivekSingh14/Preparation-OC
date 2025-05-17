package main

import "fmt"

func main11() {

	n := 3
	edges := [][]int{{0, 1}, {1, 2}, {2, 0}}
	source := 0
	destination := 2

	fmt.Println(validPath(n, edges, source, destination))

}

func validPath(n int, edges [][]int, source int, destination int) bool {

	graph := make([][]int, n)

	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	visited := make([]bool, n)

	return dfs5(graph, visited, source, destination)
}

func dfs5(graph [][]int, visited []bool, start int, end int) bool {
	if start == end {
		return true
	}

	visited[start] = true

	for _, node := range graph[start] {
		if !visited[node] && dfs5(graph, visited, node, end) {
			return true
		}
	}
	return false
}
