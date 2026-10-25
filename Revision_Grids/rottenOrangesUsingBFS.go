package main

import "fmt"

var roneighbours = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func main5() {
	grid := [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}

	cnt := 0

	cnt = rottenOrangesbfs(grid, cnt)

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				fmt.Println(-1)
				return
			}
		}
	}

	fmt.Println(cnt)

}

func rottenOrangesbfs(grids [][]int, cnt int) int {
	var que [][]int

	for i := 0; i < len(grids); i++ {
		for j := 0; j < len(grids[0]); j++ {
			if grids[i][j] == 2 {
				que = append(que, []int{i, j})
			}
		}
	}

	for len(que) != 0 {
		size := len(que)

		for k := 0; k < size; k++ {
			pair := que[0]
			que = que[1:]

			i := pair[0]
			j := pair[1]

			for k := 0; k < len(roneighbours); k++ {
				nr := i + roneighbours[k][0]
				nc := j + roneighbours[k][1]

				if nr < 0 || nc < 0 || nr >= len(grids) || nc >= len(grids[0]) || grids[nr][nc] != 1 {
					continue
				}
				grids[nr][nc] = 2
				que = append(que, []int{nr, nc})
			}
		}
		if len(que) > 0 {
			cnt++
		}

	}
	return cnt
}
