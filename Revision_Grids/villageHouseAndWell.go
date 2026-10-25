package main

import "fmt"

var vHneighbours = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func main6() {
	n := 3
	m := 3

	village := [][]byte{{'H', 'H', 'H'}, {'H', 'W', 'H'}, {'H', 'H', 'H'}}

	res := villagebfs(n, m, village)

	fmt.Println(res)

}

func villagebfs(n int, m int, village [][]byte) [][]int {

	var que [][]int

	res := make([][]int, n)

	for idx, _ := range res {
		res[idx] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if village[i][j] == 'W' {
				que = append(que, []int{i, j})
				village[i][j] = 'N'
			}
		}
	}

	for len(que) != 0 {
		pair := que[0]
		que = que[1:]

		i := pair[0]
		j := pair[1]

		for k := 0; k < len(vHneighbours); k++ {
			nr := i + vHneighbours[k][0]
			nc := j + vHneighbours[k][1]

			if nr < 0 || nc < 0 || nr >= len(village) || nc >= len(village[0]) || village[nr][nc] == 'N' {
				continue
			}
			village[nr][nc] = 'N'
			res[nr][nc] = res[i][j] + 2
			que = append(que, []int{nr, nc})
		}

	}
	return res
}
