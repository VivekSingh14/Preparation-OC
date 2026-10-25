package main

import "fmt"

var oneighbours = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func main4() {
	grid := [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}

	r := len(grid)

	c := len(grid[0])

	cnt := 0

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if grid[i][j] == 2 {
				rottenOrangesdfs(i, j, grid)
				cnt++
			}
		}
	}

	fmt.Println(cnt)

}

func rottenOrangesdfs(sr int, sc int, grids [][]int) {
	for i := 0; i < len(oneighbours); i++ {
		nr := sr + oneighbours[i][0]
		nc := sc + oneighbours[i][1]
		if nr >= 0 && nr < len(grids) && nc >= 0 && nc < len(grids[0]) && grids[nr][nc] == 1 {
			grids[nr][nc] = 2
			rottenOrangesdfs(nr, nc, grids)

		}
	}
}
