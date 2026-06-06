package main

import "fmt"

func main9() {
	V := 4
	//true case
	edges := [][]int{{0, 1}, {0, 2}, {1, 2}, {2, 3}}
	//edges := [][]int{{1, 2}, {2, 3}, {3, 1}}

	//false case
	//edges := [][]int{{0, 1}, {1, 2}, {2, 3}}
	fmt.Println(isCycle1(V, edges))
}

func isCycle1(v int, edges [][]int) bool {
	visited := make([]int, v)
	parent := make([]int, v)
	for i := 0; i < len(parent); i++ {
		parent[i] = -1
	}

	adjlist := make([][]int, v)

	for i := 0; i < len(adjlist); i++ {
		adjlist[i] = make([]int, 0)
	}

	for i := 0; i < len(edges); i++ {
		adjlist[edges[i][0]] = append(adjlist[edges[i][0]], edges[i][1])
		adjlist[edges[i][1]] = append(adjlist[edges[i][1]], edges[i][0])
	}

	for i := 0; i < v; i++ {
		if visited[i] == 0 {
			ans := bfs6(i, adjlist, visited, parent)
			if ans {
				return true
			}

		}
	}
	return false
}

func bfs6(node int, edges [][]int, visited []int, parent []int) bool {

	var queue []int

	queue = append(queue, node)
	visited[node] = 1
	for len(queue) > 0 {
		frontNode := queue[0]

		queue = queue[1:]
		for _, nbr := range edges[frontNode] {
			if visited[nbr] == 0 {
				//ans := dfs6(nbr, edges, visited, node)
				visited[nbr] = 1
				parent[nbr] = frontNode
				queue = append(queue, nbr)
			} else {
				if parent[frontNode] != nbr {
					return true
				}
			}
		}
	}
	return false
}
