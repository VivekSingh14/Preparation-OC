// given {1,2, 7}  so subset difference are {1,7} -{2} = 6 or {7,2} - {1} = 8
// or {7} - {1,2} = 4[min subset sum difference] this is we need to findout
package main

import (
	"fmt"
	"math"
)

func main8() {
	arr := []int{1, 2, 7}
	n := len(arr)
	sum := 0
	// Calculate total sum
	for _, num := range arr {
		sum += num
	}
	result := minSubsetSumDiff(arr, sum, n)
	fmt.Printf("Minimum subset sum difference: %d\n", result)
}

func minSubsetSumDiff(arr []int, sum int, n int) int {

	// Initialize 2D DP array
	target := sum / 2
	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, target+1)
	}

	// Empty subset can achieve sum 0
	for i := 0; i <= n; i++ {
		dp[i][0] = true
	}

	// Fill DP table
	for i := 1; i <= n; i++ {
		for j := 0; j <= target; j++ {
			if j < arr[i-1] {
				// Cannot include current element
				dp[i][j] = dp[i-1][j]
			} else {
				// Either include or exclude current element
				dp[i][j] = dp[i-1][j] || dp[i-1][j-arr[i-1]]
			}
		}
	}

	// Find minimum difference from last row
	minDiff := math.MaxInt32
	for j := 0; j <= target; j++ {
		if dp[n][j] {
			// If sum j is achievable, calculate difference
			s1 := j

			minDiff = min(minDiff, sum-2*s1)
		}
	}

	return minDiff
}
