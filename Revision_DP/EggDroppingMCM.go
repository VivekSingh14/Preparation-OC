package main

import (
	"fmt"
	"math"
)

func main16() {

	floor := 50
	egg := 2
	fmt.Println(solve5(egg, floor))
}

func solve5(e int, f int) int {

	if f == 0 || f == 1 {
		return f
	}

	if e == 1 {
		return f
	}

	min_num := math.MaxInt

	for k := 1; k <= f; k++ {
		temp := 1 + max(solve5(e-1, k-1), solve5(e, f-k))

		min_num = min(min_num, temp)
	}
	return min_num
}
