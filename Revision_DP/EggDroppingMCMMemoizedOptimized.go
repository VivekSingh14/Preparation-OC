package main

import (
	"fmt"
	"math"
)

var dp6 [3][1001]int

func main18() {

	floor := 100
	egg := 2

	for x := range dp6 {
		for y := range dp6[x] {
			dp6[x][y] = -1
		}
	}

	fmt.Println(solve7(egg, floor))
}

func solve7(e int, f int) int {

	if f == 0 || f == 1 {
		return f
	}

	if e == 1 {
		return f
	}

	if dp6[e][f] != -1 {
		return dp6[e][f]
	}

	min_num := math.MaxInt

	for k := 1; k <= f; k++ {
		low := 0
		high := 0
		if dp6[e-1][k-1] != -1 {
			low = dp6[e-1][k-1]
		} else {
			low = solve7(e-1, k-1)
			dp6[e-1][k-1] = low
		}

		if dp6[e][f-k] != -1 {
			high = dp6[e][f-k]
		} else {
			high = solve7(e, f-k)
			dp6[e][f-k] = high
		}

		temp := 1 + max(low, high)

		min_num = min(min_num, temp)
	}
	dp6[e][f] = min_num
	return min_num
}
