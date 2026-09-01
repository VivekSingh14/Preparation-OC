package main

import "fmt"

func main7() {

	nums := []int{9}
	target := 3

	n := len(nums)

	dp := make([]int, target+1)

	dp[0] = 1

	for i := 1; i <= target; i++ {
		for j := 0; j < n; j++ {
			if nums[j] <= i {
				dp[i] = dp[i] + dp[i-nums[j]]
			}
		}
	}

	fmt.Println(dp)

}
