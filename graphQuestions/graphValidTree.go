package main

import "fmt"

// DFS function definition
func dfs1(node, prev int, neighbors map[int][]int, visited map[int]bool) bool {
	visited[node] = true

	// If no neighbors, return false
	if len(neighbors[node]) == 0 {
		return false
	}

	// Visit all neighbors
	for _, neighbor := range neighbors[node] {
		if neighbor == prev {
			continue // Skip the previous node (to avoid going backward)
		}
		if visited[neighbor] {
			return false // Cycle detected
		}
		if !dfs1(neighbor, node, neighbors, visited) {
			return false
		}
	}

	return true
}

func validTree(n int, edges [][]int) bool {
	// Check if the number of edges is not n-1
	if len(edges) != n-1 {
		return false
	}

	// Handle the case of an empty graph (edge case)
	if len(edges) == 0 {
		return true
	}

	// Initialize neighbors map
	neighbors := make(map[int][]int)
	visited := make(map[int]bool)

	// Build the neighbors map (adjacency list)
	for _, edge := range edges {
		node1, node2 := edge[0], edge[1]
		neighbors[node1] = append(neighbors[node1], node2)
		neighbors[node2] = append(neighbors[node2], node1)
	}

	// Start DFS from node 0, with -1 as the previous node
	if !dfs1(0, -1, neighbors, visited) {
		return false
	}

	// Check if all nodes were visited
	return len(visited) == n
}

func main8() {
	// Example usage
	n := 5
	edges := [][]int{
		{0, 1},
		{0, 2},
		{0, 3},
		{1, 4},
	}

	if validTree(n, edges) {
		fmt.Println("The graph is a valid tree.")
	} else {
		fmt.Println("The graph is not a valid tree.")
	}
}
