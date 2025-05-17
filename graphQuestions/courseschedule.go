package main

import "fmt"

type void struct{}

func main9() {
	numCourses := 2
	prerequisites := [][]int{{0, 1}}

	fmt.Println(canFinish(numCourses, prerequisites))
}

func canFinish(numCourses int, prerequisites [][]int) bool {

	premap := make(map[int][]int)
	for _, pre := range prerequisites {
		premap[pre[0]] = append(premap[pre[0]], pre[1])
	}

	visited := make(map[int]void)

	for i := 0; i < numCourses; i++ {
		if !dfs2(i, visited, premap) {
			return false
		}
	}

	return true
}

func dfs2(crs int, visited map[int]void, premap map[int][]int) bool {

	// the course doesn't have any prerequisites
	if len(premap[crs]) == 0 {
		return true
	}

	if _, yes := visited[crs]; yes {
		return false // a cycle has been found because it's been visited
	}

	visited[crs] = void{}

	for _, pre := range premap[crs] {
		if !dfs2(pre, visited, premap) {
			return false
		}
	}

	delete(visited, crs)
	premap[crs] = nil

	return true
}

// func canFinish(numCourses int, prerequisites [][]int) bool {
// 	graph := map[int][]int{}
// 	for _, p := range prerequisites {
// 		graph[p[0]] = append(graph[p[0]], p[1])
// 	}
// 	visited := map[int]void{}
// 	for i := 0; i < numCourses; i++ {
// 		if !checkPrerequisites(i, visited, graph) {
// 			return false
// 		}
// 	}
// 	return true
// }

// func checkPrerequisites(course int, visited map[int]void, graph map[int][]int) bool {
// 	if len(graph[course]) == 0 {
// 		return true // the course doesn't have any prerequisites
// 	}
// 	if _, yes := visited[course]; yes {
// 		return false // a cycle has been found because it's been visited
// 	}
// 	visited[course] = void{} // being visited now
// 	for _, p := range graph[course] {
// 		if !checkPrerequisites(p, visited, graph) {
// 			return false
// 		}
// 	}
// 	// reaching this point means no cycles
// 	delete(visited, course)
// 	graph[course] = nil // a course that can be completed is the same as having no prerequisites.
// 	return true
// }
