package main

import "fmt"

// edge list is given:
// {1,2}, {2,3}, {3,4}, {4,2}, {1,3}
func main2() {

	edgelist := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 2}, {1, 3}}
	//number of edges given or we need to  findout in question
	n := 5

	adjmatrix := make([][]int, n)

	for i := 0; i < len(adjmatrix); i++ {
		adjmatrix[i] = make([]int, n)
	}

	for i := 0; i < len(edgelist); i++ {
		a := edgelist[i][0]
		b := edgelist[i][1]
		adjmatrix[a][b] = 1
		adjmatrix[b][a] = 1

	}

	traverseGraph1(adjmatrix)

}

func traverseGraph1(adjmatrix [][]int) {

	for i := 1; i < len(adjmatrix); i++ {
		fmt.Print("Node: ", i, " Neighbours: ")
		for j := 1; j < len(adjmatrix[i]); j++ {
			if adjmatrix[i][j] == 1 {
				fmt.Print(j, " ")
			}
		}
		fmt.Println()
	}
}
