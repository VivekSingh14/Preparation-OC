package main

import "fmt"

// edge list is given:
// {1,2}, {2,3}, {3,4}, {4,2}, {1,3}

// Representation of graph:
//
//	1 - 2---|
//	|   |   |
//	3 ---   |
//	V       |
//	4-------
//
// BFS :   1,2,3,4
type Graph1 struct {
	adjlist map[int][]int
}

func NewGraph1() *Graph1 {
	return &Graph1{adjlist: make(map[int][]int)}
}

func (g *Graph1) addEdges(u int, v int) {
	g.adjlist[u] = append(g.adjlist[u], v)
	//if graph is undirected
	g.adjlist[v] = append(g.adjlist[v], u)
}

func traverseGraph2(g *Graph1) {

	for u, v := range g.adjlist {

		fmt.Println("vertex: ", u, "and its neighbours are: ", v)
	}

}

func main3() {

	g := NewGraph1()

	g.addEdges(1, 2)
	g.addEdges(2, 3)
	g.addEdges(3, 4)
	g.addEdges(4, 2)
	g.addEdges(1, 3)

	//traverseGraph2(g)

	source := 1
	BFS(source, g)

}

func BFS(source int, g *Graph1) {
	var queue []int

	queue = append(queue, source)

	visited := make([]int, 5)

	// for len(queue) != 0 {

	// 	if val, ok := g.adjlist[queue[0]]; ok {
	// 		fmt.Print(queue[0], " ")

	// 		for _, node := range val {
	// 			if visited[node] == 0 {
	// 				queue = append(queue, node)
	// 			}
	// 		}
	// 		visited[queue[0]] = 1
	// 		queue = queue[1:]
	// 	}
	// }

	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]

		fmt.Print(node, " ")

		visited[node] = 1

		if val, ok := g.adjlist[node]; ok {
			for _, neighbor := range val {
				if visited[neighbor] == 0 {
					queue = append(queue, neighbor)
					visited[neighbor] = 1 // Mark as visited when enqueued
				}
			}
		}
	}

}
