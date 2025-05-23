package main

import "fmt"

func main() {

	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0

	fmt.Println(search(nums, target))
}

func search(nums []int, target int) int {
	// left, right := 0, len(nums)-1

	// for left < right {
	// 	mid := (left + right) / 2

	// 	if nums[left] > nums[mid] {
	// 		if nums[mid] < target && target <= nums[right] {
	// 			left = mid + 1
	// 		} else {
	// 			right = mid - 1
	// 		}
	// 	} else {
	// 		if nums[left] <= target && target < nums[mid] {
	// 			right = mid - 1
	// 		} else {
	// 			left = mid + 1
	// 		}
	// 	}
	// }
	// return -1

	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}

		if nums[left] > nums[mid] {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}

	return -1
}
