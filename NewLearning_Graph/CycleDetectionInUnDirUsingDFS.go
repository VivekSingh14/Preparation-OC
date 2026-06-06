package main

import (
	"fmt"
	"log"
)

var temp_count int = 0

func main18() {
	V := 4
	//true case
	edges := [][]int{{0, 1}, {0, 2}, {1, 2}, {2, 3}}
	//edges := [][]int{{1, 2}, {2, 3}, {3, 1}}

	//false case
	//edges := [][]int{{0, 1}, {1, 2}, {2, 3}}
	fmt.Println(isCycle(V, edges))
}

func isCycle(v int, edges [][]int) bool {
	visited := make([]int, v)

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
			ans := dfs6(i, adjlist, visited, -1)
			if ans {
				return true
			}

		}
	}
	return false
}

func dfs6(node int, edges [][]int, visited []int, parent int) bool {
	temp_count++
	log.Printf("node %d  parent  %d  countdfs %d  \n", node, parent, temp_count)
	visited[node] = 1

	for _, nbr := range edges[node] {
		log.Printf("nbr %d \n", nbr)
		if visited[nbr] == 0 {
			ans := dfs6(nbr, edges, visited, node)
			if ans {
				return true
			}
		} else if visited[nbr] == 1 && nbr != parent {
			return true
		}
	}
	return false
}
