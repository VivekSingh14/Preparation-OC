package main

import "fmt"

func main4() {
	heights := [][]int{{1, 2, 2, 3, 5}, {3, 2, 3, 4, 4}, {2, 4, 5, 3, 1}, {6, 7, 1, 4, 5}, {5, 1, 1, 2, 4}}
	fmt.Println(pacificAtlantic(heights))

}

func pacificAtlantic(heights [][]int) [][]int {
	res := make([][]int, 0)
	if len(heights) == 0 || len(heights[0]) == 0 {
		return res
	}

	m := len(heights)
	n := len(heights[0])
	pacific := make([][]bool, m)
	for index, _ := range pacific {
		pacific[index] = make([]bool, n)
	}
	atlantic := make([][]bool, m)
	for index, _ := range atlantic {
		atlantic[index] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		dfs(heights, pacific, -1, i, 0)
		dfs(heights, atlantic, -1, i, n-1)
	}

	for i := 0; i < n; i++ {
		dfs(heights, pacific, -1, 0, i)
		dfs(heights, atlantic, -1, m-1, i)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if pacific[i][j] && atlantic[i][j] {
				res = append(res, []int{i, j})
			}
		}
	}

	return res
}

func dfs(heights [][]int, visit [][]bool, pre, i, j int) {
	m := len(heights)
	n := len(heights[0])

	if i < 0 || i >= m || j < 0 || j >= n || visit[i][j] || pre > heights[i][j] {
		return
	}

	visit[i][j] = true

	dfs(heights, visit, heights[i][j], i, j+1)
	dfs(heights, visit, heights[i][j], i, j-1)
	dfs(heights, visit, heights[i][j], i+1, j)
	dfs(heights, visit, heights[i][j], i-1, j)
}
