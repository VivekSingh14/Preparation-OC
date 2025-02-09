package main

import (
	"fmt"
	"math"
)

func main1() {
	//nums := []int{4, 5, 1, 2, 3}
	nums := []int{4, 5, 6, 7, 0, 1, 2}

	fmt.Println(findMin(nums))
	fmt.Println(findMinBS(nums))
}

func findMin(nums []int) int {

	if len(nums) <= 1 {
		return nums[0]
	}

	low := 0
	high := len(nums) - 1

	min := math.MaxInt

	for low <= high {

		if nums[low] < min {
			min = nums[low]
		} else if nums[high] < min {
			min = nums[high]
		} else {
			low++
			high--
		}
	}

	return min
}

func findMinBS(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := (left + right) / 2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return nums[left]
}
