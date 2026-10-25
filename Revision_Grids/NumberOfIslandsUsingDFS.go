package main

import "fmt"

var gneighbours = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func main2() {
	grids := [][]int{{1, 1, 1, 1, 0}, {1, 1, 0, 1, 0}, {1, 1, 0, 0, 0}, {0, 0, 0, 0, 0}}

	r := len(grids)

	c := len(grids[0])

	cnt := 0

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if grids[i][j] == 1 {
				islanddfs(i, j, grids)
				cnt++
			}
		}
	}

	fmt.Println(cnt)

}

func islanddfs(sr int, sc int, grids [][]int) {
	grids[sr][sc] = 2
	for i := 0; i < len(gneighbours); i++ {
		nr := sr + gneighbours[i][0]
		nc := sc + gneighbours[i][1]
		if nr >= 0 && nr < len(grids) && nc >= 0 && nc < len(grids[0]) && grids[nr][nc] == 1 {

			islanddfs(nr, nc, grids)
		}
	}
}
