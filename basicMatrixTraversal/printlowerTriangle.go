package main

import "fmt"

func main() {
	matrix := make([][]int, 5)

	for index := range matrix {
		matrix[index] = make([]int, 5)
	}

	count := 0

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			matrix[i][j] = count
			count++
		}
	}

	// lower triangle
	for i := 0; i < len(matrix); i++ {
		for j := 0; j <= i; j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}

	fmt.Println("------------------")

	// upper triangle
	for i := 0; i < len(matrix); i++ {

		for j := 0; j < i; j++ {
			fmt.Print(" ")
		}
		for j := i; j < len(matrix[i]); j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}

	fmt.Println("------------------")

	// print mid column
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i][len(matrix[i])/2])
	}

	fmt.Println("------------------")

	// print mid row

	for i := 0; i < len(matrix); i++ {
		fmt.Print(matrix[len(matrix)/2][i], " ")
	}
	fmt.Println()

}
