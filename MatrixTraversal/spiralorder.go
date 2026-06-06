package main

import "fmt"

func main3() {

	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	fmt.Println(spiralOrder(matrix))
}

func spiralOrder(matrix [][]int) []int {

	top := 0
	left := 0
	right := len(matrix[0]) - 1
	down := len(matrix) - 1

	var elements []int

	for left <= right && top <= down {

		for i := left; i <= right; i++ {
			elements = append(elements, matrix[top][i])
		}
		top++

		for i := top; i <= down; i++ {
			elements = append(elements, matrix[i][right])
		}
		right--

		for i := right; left <= i; i-- {
			elements = append(elements, matrix[down][i])
		}
		down--

		for i := down; top <= i; i-- {
			elements = append(elements, matrix[i][left])
		}
		left++
	}

	area := len(matrix) * len(matrix[0])
	return elements[:area]

}
