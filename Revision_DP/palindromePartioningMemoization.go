package main

import (
	"fmt"
	"math"
)

var dp7 [2001][2001]int

func main22() {
	str1 := "aab"

	for i := range dp7 {
		for j := range dp7[i] {
			dp7[i][j] = -1
		}
	}

	fmt.Println(solvepalin1(str1, 0, len(str1)-1))
}

func solvepalin1(str string, i int, j int) int {
	if i >= j {
		return 0
	}

	if isPalindrome1(str, i, j) {
		return 0
	}

	if dp7[i][j] != -1 {
		return dp7[i][j]
	}
	min_num := math.MaxInt

	for k := i; k <= j-1; k++ {
		temp1 := 0
		temp2 := 0
		if dp7[i][k] != -1 {
			temp1 = dp7[i][k]
		} else {
			temp1 = solvepalin1(str, i, k)
		}
		if dp7[k+1][j] != -1 {
			temp2 = dp7[k+1][j]
		} else {
			temp2 = solvepalin1(str, k+1, j)
		}
		temp := 1 + temp1 + temp2
		if temp < min_num {
			min_num = temp
		}
	}
	dp7[i][j] = min_num
	return min_num
}

func isPalindrome1(str string, i int, j int) bool {

	for i <= j {
		if str[i] != str[j] {
			return false
		}
		i++
		j--
	}
	return true
}
