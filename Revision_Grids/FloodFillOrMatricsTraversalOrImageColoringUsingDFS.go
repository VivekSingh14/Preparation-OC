package main

import "fmt"

var neighbours = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func main1() {

	image := [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}

	visited := make([][]int, len(image))

	for idx, _ := range visited {
		visited[idx] = make([]int, len(image[0]))
	}

	sr := 1
	sc := 1
	color := 2

	dfs(sr, sc, image, visited, color)

	fmt.Println(image)

}

func dfs(sr int, sc int, image [][]int, visited [][]int, color int) {
	visited[sr][sc] = 1
	originalcolor := image[sr][sc]
	image[sr][sc] = color

	for i := 0; i < len(neighbours); i++ {
		nr := sr + neighbours[i][0]
		nc := sc + neighbours[i][1]
		if nr >= 0 && nr < len(image) && nc >= 0 && nc < len(image[0]) && visited[nr][nc] == 0 && image[nr][nc] == originalcolor {

			dfs(nr, nc, image, visited, color)
		}
	}
}
