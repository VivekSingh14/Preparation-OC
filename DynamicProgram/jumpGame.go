package main

import "fmt"

func main8() {

	nums := []int{3, 2, 1, 0, 4}

	fmt.Println(canJump(nums))
}

func canJump(nums []int) bool {

	/* Method-1: Greedy */

	goal := len(nums) - 1

	for i := len(nums) - 1; i >= 0; i-- {
		if i+nums[i] >= goal {
			goal = i
		}
	}

	return goal == 0

	/* Method 2 */
	// farthest := 0
	// n := len(nums)

	// for i := 0; i < n; i++ {
	// 	if i+nums[i] > farthest {
	// 		farthest = i + nums[i]
	// 	}

	// 	if nums[i] == 0 && i < n-1 && i == farthest {
	// 		return false
	// 	}
	// }

	// return true
}
