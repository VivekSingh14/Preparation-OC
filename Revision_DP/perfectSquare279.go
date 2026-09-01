package main

import (
	"fmt"
	"math"
)

func main6() {
	n := 13

	nums := []int{}

	for i := 1; i*i <= n; i++ {
		nums = append(nums, i*i)
	}

	m := len(nums)

	dp := make([][]int, m+1)

	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i <= m; i++ {
		dp[i][0] = 0
	}

	for i := 0; i <= n; i++ {
		dp[0][i] = math.MaxInt - 1
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if nums[i-1] <= j {
				dp[i][j] = min(dp[i-1][j], 1+dp[i][j-nums[i-1]])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	fmt.Println(dp[m][n])
}
