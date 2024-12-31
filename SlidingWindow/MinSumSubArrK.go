package main

import (
	"fmt"
	"math"
)

// ## Min Sum SubArray with Size K
// ## Sliding Window Algo
func main1() {
	arr := []int{10, 4, 2, 5, 6, 3, 8, 1}
	k := 3

	res, err := minSumSubArray(arr, k)
	if err != nil {
		fmt.Println("Wrong input")
	}

	fmt.Println(arr[res[0]:res[1]])
}

func minSumSubArray(arr []int, k int) (res []int, err error) {
	min_window := math.MaxInt
	window_sum := 0
	for i := 0; i < k; i++ {
		window_sum += arr[i]
	}
	min_window = window_sum
	start := 0
	end := k - 1
	for i := k; i < len(arr); i++ {
		window_sum = window_sum + arr[i] - arr[i-k]
		if window_sum < min_window {
			min_window = window_sum
			start = i - k
			end = i
		}
	}

	res = append(res, start)
	res = append(res, end)

	return res, nil
}
