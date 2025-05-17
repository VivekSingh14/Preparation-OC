package main

import "fmt"

func main14() {
	edges := [][]int{{1, 2}, {2, 3}, {4, 2}}
	fmt.Println(findCenter(edges))
}

func findCenter(edges [][]int) int {

	graph := make(map[int][]int)

	for _, edge := range edges {
		node1, node2 := edge[0], edge[1]
		graph[node1] = append(graph[node1], node2)
		graph[node2] = append(graph[node2], node1)
	}

	for key, val := range graph {
		if len(val) == len(edges) {
			return key
		}
	}
	return -1
}
