package main

import "fmt"

type Graph2 struct {
	adjlist map[int][]int
}

func NewGraph2() *Graph2 {
	return &Graph2{adjlist: make(map[int][]int)}
}

func (g *Graph2) appendalledges(u int, v int) {
	g.adjlist[u] = append(g.adjlist[u], v)
	g.adjlist[v] = append(g.adjlist[v], u)
}

func bfsIfPathExists(g *Graph2, source int, visited []int) {
	var que []int
	que = append(que, source)
	visited[source] = 1

	for len(que) != 0 {
		node := que[0]
		que = que[1:]

		if val, ok := g.adjlist[node]; ok {
			for _, neighbour := range val {
				if visited[neighbour] == 0 {
					que = append(que, neighbour)
					visited[neighbour] = 1
				}
			}
		}

	}
}

func main7() {

	edges := [][]int{{0, 1}, {1, 2}, {2, 0}}
	n := 3
	source := 0
	destination := 2

	g := NewGraph2()
	for _, v := range edges {
		g.appendalledges(v[0], v[1])
	}

	visited := make([]int, n)

	bfsIfPathExists(g, source, visited)

	if visited[destination] == 1 {
		fmt.Println(visited[destination] == 1)
	}

}
