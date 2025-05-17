package main

import "fmt"

func main12() {

	// true case

	// edges := [][]int{{0, 1}, {1, 2}, {2, 0}}
	// source := 0
	// destination := 2

	// false case

	edges := [][]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}}
	source := 0
	destination := 5

	fmt.Println(validPath1(edges, source, destination))
}

func validPath1(edges [][]int, source, destination int) bool {

	graph := make(map[int][]int)
	visited := make(map[int]bool)

	for _, edge := range edges {
		node1, node2 := edge[0], edge[1]
		graph[node1] = append(graph[node1], node2)
		graph[node2] = append(graph[node2], node1)
	}

	return dfs9(graph, visited, source, destination)

}

func dfs9(graph map[int][]int, visited map[int]bool, source, destination int) bool {
	if source == destination {
		return true
	}

	visited[source] = true

	for _, node := range graph[source] {
		if !visited[node] && dfs9(graph, visited, node, destination) {
			return true
		}
	}
	return false
}
