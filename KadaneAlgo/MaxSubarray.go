package main

import (
	"fmt"
	"math"
)

func main() {

	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	fmt.Println(maxSubArray(arr))

}

func maxSubArray(arr []int) int {

	max_so_far := math.MinInt

	max_end_here := 0

	if len(arr) <= 1 {
		return arr[0]
	}

	for i := 0; i < len(arr); i++ {
		max_end_here += arr[i]
		if max_so_far < max_end_here {
			max_so_far = max_end_here
		}
		if max_end_here < 0 {
			max_end_here = 0
		}
	}
	return max_so_far
}
