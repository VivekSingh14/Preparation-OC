package main

import "fmt"

func main() {
	numCourses := 2
	prerequisites := [][]int{{1, 0}}

	fmt.Println(canFinish(numCourses, prerequisites))
}

func canFinish(numCourses int, prerequisites [][]int) bool {

	premap := make(map[int][][]int)
	for crs, pre := range prerequisites {
		var temp [][]int
		temp = append(temp, pre)
		premap[crs] = temp
	}

	visited := make(map[int]bool)

	for i := 0; i < numCourses; i++ {
		dfs2(i, visited, premap)
	}

	return true
}

func dfs2(crs int, visited map[int]bool, premap map[int][][]int) bool {

	if visited[crs] {
		return false
	}

	if len(premap[crs]) == 0 {
		return true
	}

	visited[crs] = true

	for pre, _ := range premap[crs] {
		if !dfs2(pre, visited, premap) {
			return false
		}
	}

	return false
}
