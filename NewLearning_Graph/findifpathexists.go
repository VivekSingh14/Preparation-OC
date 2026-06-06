package main

import "fmt"

// edge list is given:
// {1,2}, {2,3}, {3,4}, {4,2}, {1,3}

// Reprecentation of graph:
//
//	1 - 2---|
//	|   |   |
//	3 ---   |
//	V       |
//	4-------
//
// BFS :   1, 2,3,4
type GraphDFS1 struct {
	adjlist map[int][]int
}

func NewGraphDFS1() *GraphDFS1 {
	return &GraphDFS1{adjlist: make(map[int][]int)}
}

func (g *GraphDFS1) addEdges(u int, v int) {
	g.adjlist[u] = append(g.adjlist[u], v)
	//if graph is undirected
	g.adjlist[v] = append(g.adjlist[v], u)
}

func traverseGraph4(g *GraphDFS1) {

	for u, v := range g.adjlist {

		fmt.Println("vertex: ", u, "and its neighbours are: ", v)
	}

}

func main5() {

	g := NewGraphDFS1()

	g.addEdges(1, 2)
	g.addEdges(2, 3)
	g.addEdges(3, 4)
	g.addEdges(4, 2)
	g.addEdges(1, 3)

	// g.addEdges(0, 1)
	// g.addEdges(1, 2)
	// g.addEdges(2, 0)

	//traverseGraph3(g)

	source := 1
	destination := 4
	visited := make([]int, 5)
	DFS1(source, g, visited)

	fmt.Println(visited[destination] == 1)

}

func DFS1(source int, g *GraphDFS1, visited []int) {

	fmt.Println(source)
	visited[source] = 1

	if val, ok := g.adjlist[source]; ok {
		for _, neighbor := range val {
			if visited[neighbor] == 0 {
				DFS1(neighbor, g, visited)
			}
		}
	}
}
