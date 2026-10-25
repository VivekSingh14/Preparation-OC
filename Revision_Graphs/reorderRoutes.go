package main

import "fmt"

func main9() {
	n := 6
	connections := [][]int{{0, 1}, {1, 3}, {2, 3}, {4, 0}, {4, 5}}

	fmt.Println(minOrder(n, connections))

}

func minOrder(n int, connections [][]int) int {
	forward_nbrs := make(map[int][]int)
	backward_nbrs := make(map[int][]int)

	visited := make([]int, n)

	for i := 0; i < len(connections); i++ {
		a := connections[i][0]
		b := connections[i][1]

		forward_nbrs[a] = append(forward_nbrs[a], b)
		backward_nbrs[b] = append(backward_nbrs[b], a)
	}
	ans := 0
	reorderDFS(0, forward_nbrs, backward_nbrs, &ans, visited)
	return ans

}

func reorderDFS(source int, forward_nbrs map[int][]int, backward_nbrs map[int][]int, ans *int, visited []int) {

	visited[source] = 1

	if nbrs, ok := forward_nbrs[source]; ok {
		for _, nbr := range nbrs {
			if visited[nbr] == 0 {
				(*ans)++
				reorderDFS(nbr, forward_nbrs, backward_nbrs, ans, visited)
			}
		}
	}

	if nbrs, ok := backward_nbrs[source]; ok {
		for _, nbr := range nbrs {
			if visited[nbr] == 0 {
				reorderDFS(nbr, forward_nbrs, backward_nbrs, ans, visited)
			}
		}
	}

}
