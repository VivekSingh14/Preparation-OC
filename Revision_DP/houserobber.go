package main

import "fmt"

func main28() {
	num := []int{2, 7, 9, 3, 1}

	fmt.Println(rob(num))
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums)+1)
	dp[1] = nums[0]
	dp[2] = max(nums[0], nums[1])
	for i := 3; i <= len(nums); i++ {
		dp[i] = max(dp[i-2]+nums[i-1], dp[i-1])
	}

	return dp[len(nums)]
}
