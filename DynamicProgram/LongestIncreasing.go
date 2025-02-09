package main

import (
	"fmt"
	"math"
)

func main4() {

	nums := []int{0, 1, 0, 3, 2, 3}

	fmt.Println(lengthOfLIS(nums))
}

func lengthOfLIS(nums []int) int {

	lis := make([]int, len(nums))

	for i := 0; i < len(lis); i++ {
		lis[i] = 1
	}

	for i := len(nums) - 1; i >= 0; i-- {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] < nums[j] {
				lis[i] = max(lis[i], 1+lis[j])
			}
		}
	}
	temp := math.MinInt

	for i := 0; i < len(lis); i++ {
		if lis[i] > temp {
			temp = lis[i]
		}
	}
	return temp
}
