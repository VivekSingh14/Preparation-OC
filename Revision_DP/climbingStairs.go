package main

import "fmt"

func main27() {

	num := 3

	fmt.Println(climbStairs(num))
}
func climbStairs(n int) int {
	if n <= 1 {
		return n
	}

	dp := make([]int, n+1)

	dp[1] = 1
	dp[2] = 2

	if dp[n] > 0 {
		return dp[n]
	}

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}
