package main

import (
	"fmt"
	"math"
)

func main6() {

	//graph = [][]int {{0,1,5},{1,2,3},{2,3,1},{0,3,10}}

	matrix := [][]int{{0, 5, math.MaxInt, 10}, {math.MaxInt, 0, 3, math.MaxInt}, {math.MaxInt, math.MaxInt, 0, 1}, {math.MaxInt, math.MaxInt, math.MaxInt, 0}}

	v := 4

	for k := 0; k < v; k++ {
		for i := 0; i < v; i++ {
			for j := 0; j < v; j++ {
				if matrix[i][k] == math.MaxInt || matrix[k][j] == math.MaxInt {
					continue
				}

				dist_through_k := matrix[i][k] + matrix[k][j]
				if dist_through_k < matrix[i][j] {
					matrix[i][j] = dist_through_k
				}
			}
		}
	}

	fmt.Println(matrix)

}
