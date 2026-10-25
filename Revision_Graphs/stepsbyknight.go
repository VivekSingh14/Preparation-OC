package main

import "fmt"

var coordinates = [][]int{{-1, -2}, {-1, 2}, {1, -2}, {1, 2}, {-2, -1}, {-2, 1}, {2, -1}, {2, 1}}

func main8() {
	n := 6
	src := []int{4, 5}
	target := []int{1, 1}

	src_x := n - src[1]
	src_y := src[0] - 1

	target_x := n - target[1]
	target_y := target[0] - 1

	fmt.Println(findSteps(src_x, src_y, target_x, target_y, n))
}

func findSteps(src_x int, src_y int, target_x int, target_y int, n int) int {
	visited := make([][]int, n)

	for i := 0; i < n; i++ {
		visited[i] = make([]int, n)
	}
	var que [][]int
	que = append(que, []int{src_x, src_y})

	visited[src_x][src_y] = 1

	steps := 0

	for len(que) != 0 {
		qSize := len(que)

		for qSize > 0 {
			step := que[0]
			que = que[1:]

			if step[0] == target_x && step[1] == target_y {
				return steps
			}

			for k := 0; k < len(coordinates); k++ {
				new_x := step[0] + coordinates[k][0]
				new_y := step[1] + coordinates[k][1]

				if new_x >= 0 && new_y >= 0 && new_x < n && new_y < n && visited[new_x][new_y] == 0 {
					visited[new_x][new_y] = 1
					que = append(que, []int{new_x, new_y})
				}
			}
			qSize--
		}
		steps += 1
	}
	return -1
}
