package main

// Valid Path using DFS
import (
	"fmt"
)

// Unweighted Bidirectional graph
type Graph1 struct {
	adjlist map[int][]int
}

func NewGraph1() *Graph1 {
	return &Graph1{adjlist: make(map[int][]int)}
}

func (g *Graph1) Append1(u int, v int) {
	g.adjlist[u] = append(g.adjlist[u], v)
	g.adjlist[v] = append(g.adjlist[v], u)
}

func DFS1(s *Graph1, source int, visited []int) {
	fmt.Println(source)
	visited[source] = 1

	if val, ok := s.adjlist[source]; ok {
		for _, neighbour := range val {
			if visited[neighbour] == 0 {
				DFS1(s, neighbour, visited)
			}
		}
	}
}

// most of the logic or functionality should be moved to validPath function from main and Append1 function
func main6() {

	edges := [][]int{{0, 1}, {1, 2}, {2, 0}}
	n := 3
	source := 0
	destination := 2

	g1 := NewGraph1()
	for _, v := range edges {
		g1.Append1(v[0], v[1])
	}

	visited := make([]int, n)

	DFS1(g1, source, visited)

	if visited[destination] == 1 {
		fmt.Println(visited[destination] == 1)
	}
}
