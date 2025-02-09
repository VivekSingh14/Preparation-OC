package main

import "fmt"

func main2() {

	fmt.Println(climbStairs(1))

}

func climbStairs(n int) int {

	dp := make([]int, n+1)

	dp[n] = 1
	dp[n-1] = 1

	for i := n - 2; i >= 0; i-- {
		dp[i] = dp[i+1] + dp[i+2]
	}
	return dp[0]
}
