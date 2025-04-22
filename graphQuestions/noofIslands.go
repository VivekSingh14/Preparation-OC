package main

import "fmt"

func main1() {
	grid := [][]byte{{1, 1, 1, 1, 0}, {1, 1, 0, 1, 0}, {1, 1, 0, 0, 0}, {0, 0, 0, 0, 0}}

	fmt.Println(numIslands(grid))

}

func numIslands(grid [][]byte) int {
	islands := 0

	rows, cols := len(grid), len(grid[0])

	type pair struct {
		r int
		c int
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 {
				queue := []pair{}
				queue = append(queue, pair{r: i, c: j})
				grid[i][j] = 0

				for len(queue) > 0 {
					currentPair := queue[0]
					queue = queue[1:]

					row, col := currentPair.r, currentPair.c

					//up
					if row-1 >= 0 && grid[row-1][col] == 1 {
						queue = append(queue, pair{r: row - 1, c: col})
						grid[row-1][col] = 0
					}

					//down
					if row+1 < rows && grid[row+1][col] == 1 {
						queue = append(queue, pair{r: row + 1, c: col})
						grid[row+1][col] = 0
					}

					//left
					if col-1 >= 0 && grid[row][col-1] == 1 {
						queue = append(queue, pair{r: row, c: col - 1})
						grid[row][col-1] = 0
					}

					//right
					if col+1 <= cols && grid[row][col+1] == 1 {
						queue = append(queue, pair{r: row, c: col + 1})
						grid[row][col+1] = 0
					}

				}
				islands++
			}
		}
	}
	return islands
}
