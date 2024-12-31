package main

import "fmt"

func main() {
	grid := [][]byte{{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'}}
	fmt.Println(numIslands(grid))

}

func numIslands(grid [][]byte) int {
	result := 0
	for i, row := range grid {
		for j, value := range row {
			if value == '1' {
				visit(&grid, i, j)
				result += 1
			}
		}
	}
	return result
}

func visit(grid *[][]byte, i int, j int) {
	m, n := len(*grid), len((*grid)[0])
	if (i < 0) || (i >= m) || (j < 0) || (j >= n) || ((*grid)[i][j] == '0') {
		return
	}

	(*grid)[i][j] = '0'

	visit(grid, i-1, j)
	visit(grid, i+1, j)
	visit(grid, i, j-1)
	visit(grid, i, j+1)
}
