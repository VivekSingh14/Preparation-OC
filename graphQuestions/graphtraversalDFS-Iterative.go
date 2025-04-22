package main

import "fmt"

type Graph struct {
	adjList map[int][]int
}

func NewGraph() *Graph {
	return &Graph{adjList: make(map[int][]int)}
}

func (g *Graph) AddEdge(u, v int) {
	g.adjList[u] = append(g.adjList[u], v)
	g.adjList[v] = append(g.adjList[v], u)
}

func (g *Graph) DFS(start int) {
	visited := make(map[int]bool)

	stack := []int{start}

	for len(stack) > 0 {

		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if visited[node] {
			continue
		}

		visited[node] = true

		fmt.Print(node, " ")

		for _, neighbor := range g.adjList[node] {

			if !visited[neighbor] {
				stack = append(stack, neighbor)
			}
		}
	}

	fmt.Println()

}

func main6() {

	graph := NewGraph()

	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(1, 4)
	graph.AddEdge(2, 5)

	// for u, v := range graph.adjList {

	// 	fmt.Println("Key: ", u, "  Value: ", v)
	// }

	graph.DFS(0)

	graph.DFS(1)

}
