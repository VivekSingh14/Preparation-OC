package main

import (
	"fmt"
	"math"
)

func main21() {
	str1 := "aab"

	fmt.Println(solvepalin(str1, 0, len(str1)-1))
}

func solvepalin(str string, i int, j int) int {
	if i >= j {
		return 0
	}

	if isPalindrome(str, i, j) {
		return 0
	}
	min_num := math.MaxInt

	for k := i; k <= j-1; k++ {
		temp := 1 + solvepalin(str, i, k) + solvepalin(str, k+1, j)
		if temp < min_num {
			min_num = temp
		}
	}
	return min_num
}

func isPalindrome(str string, i int, j int) bool {

	for i <= j {
		if str[i] != str[j] {
			return false
		}
		i++
		j--
	}
	return true
}
