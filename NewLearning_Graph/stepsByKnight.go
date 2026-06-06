package main

import "fmt"

var coordinates [][]int = [][]int{{-1, -2}, {-1, 2}, {1, -2}, {1, 2}, {-2, -1}, {-2, 1}, {2, -1}, {2, 1}}

type node struct {
	x, y int
}

func main6() {
	n := 6
	knightposition := []int{1, 3}
	targetposition := []int{5, 0}

	fmt.Println(minstepstoreachtarget(knightposition, targetposition, n))
}

func minstepstoreachtarget(knightposition []int, targetposition []int, n int) int {
	// if the knight and target position is given in the form of chess blocks or x-y graph form
	// src_x := n-knightposition[1] //(n-y)
	// src_y := knightposition[0]-1 //(x-1)
	// target_x := n-targetposition[1]
	// target_y := targetposition[0]-1

	return helper(knightposition[0], knightposition[1], targetposition[0], targetposition[1], n)

}

func helper(src_x int, src_y int, target_x int, target_y int, n int) int {
	visited := make([][]int, n)

	for i := 0; i < n; i++ {
		visited[i] = make([]int, n)
	}

	var queue []node

	queue = append(queue, node{src_x, src_y})
	visited[src_x][src_y] = 1

	steps := 0

	for len(queue) != 0 {
		tempqSize := len(queue)
		for tempqSize > 0 {
			temp := queue[0]
			queue = queue[1:]
			if temp.x == target_x && temp.y == target_y {
				return steps
			}
			for i := 0; i < len(coordinates); i++ {
				new_x := temp.x + coordinates[i][0]
				new_y := temp.y + coordinates[i][1]
				if new_x >= 0 && new_y >= 0 && new_x < n && new_y < n && visited[new_x][new_y] == 0 {
					visited[new_x][new_y] = 1
					queue = append(queue, node{new_x, new_y})
				}
			}
			tempqSize--
		}
		steps++
	}
	return -2
}
