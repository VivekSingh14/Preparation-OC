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
// BFS :   1,2,3,4
type GraphDFS struct {
	adjlist map[int][]int
}

func NewGraphDFS() *GraphDFS {
	return &GraphDFS{adjlist: make(map[int][]int)}
}

func (g *GraphDFS) addEdges(u int, v int) {
	g.adjlist[u] = append(g.adjlist[u], v)
	//if graph is undirected
	g.adjlist[v] = append(g.adjlist[v], u)
}

func traverseGraph3(g *GraphDFS) {

	for u, v := range g.adjlist {

		fmt.Println("vertex: ", u, "and its neighbours are: ", v)
	}

}

func main4() {
	// Chatgpt prompt
	// any easy programming question on kruskal's mst algo...must be in goland and explain it step by step
	g := NewGraphDFS()

	g.addEdges(1, 2)
	g.addEdges(2, 3)
	g.addEdges(3, 4)
	g.addEdges(4, 2)
	g.addEdges(1, 3)

	//traverseGraph3(g)

	source := 1
	visited := make([]int, 5)
	DFS(source, g, visited)

}

func DFS(source int, g *GraphDFS, visited []int) {

	fmt.Println(source)
	visited[source] = 1

	if val, ok := g.adjlist[source]; ok {
		for _, neighbor := range val {
			if visited[neighbor] == 0 {
				DFS(neighbor, g, visited)
			}
		}
	}
}

// package main

// import "fmt"

// // edge list is given:
// // {1,2}, {2,3}, {3,4}, {4,2}, {1,3}

// // Reprecentation of graph:
// //
// //	1 - 2---|
// //	|   |   |
// //	3 ---   |
// //	V       |
// //	4-------
// //
// // BFS :   1, 2,3,4
// type GraphDFS struct {
// 	adjlist map[int][]int
// }

// func NewGraphDFS() *GraphDFS {
// 	return &GraphDFS{adjlist: make(map[int][]int)}
// }

// func (g *GraphDFS) addEdges(u int, v int) {
// 	g.adjlist[u] = append(g.adjlist[u], v)
// 	//if graph is undirected
// 	g.adjlist[v] = append(g.adjlist[v], u)
// }

// func traverseGraph3(g *GraphDFS) {

// 	for u, v := range g.adjlist {

// 		fmt.Println("vertex: ", u, "and its neighbours are: ", v)
// 	}

// }

// func main() {

// 	g := NewGraphDFS()

// 	g.addEdges(1, 2)
// 	g.addEdges(2, 3)
// 	g.addEdges(3, 4)
// 	g.addEdges(4, 2)
// 	g.addEdges(1, 3)

// 	//traverseGraph3(g)

// 	source := 1
// 	DFS(source, g)

// }

// func DFS(source int, g *GraphDFS) {

// }
