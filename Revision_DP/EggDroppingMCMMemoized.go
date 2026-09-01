package main

import (
	"fmt"
	"math"
)

var dp5 [3][1001]int

func main17() {

	floor := 100
	egg := 2

	for x := range dp5 {
		for y := range dp5[x] {
			dp5[x][y] = -1
		}
	}

	fmt.Println(solve6(egg, floor))
}

func solve6(e int, f int) int {

	if f == 0 || f == 1 {
		return f
	}

	if e == 1 {
		return f
	}

	if dp5[e][f] != -1 {
		return dp5[e][f]
	}

	min_num := math.MaxInt

	for k := 1; k <= f; k++ {
		temp := 1 + max(solve6(e-1, k-1), solve6(e, f-k))

		min_num = min(min_num, temp)
	}
	dp5[e][f] = min_num
	return min_num
}
