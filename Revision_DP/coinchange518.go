package main

import (
	"fmt"
	"math"
)

func main5() {
	coins := []int{2, 5}
	amount := 5

	n := len(coins)

	dp := make([][]int, n+1)

	for i := 0; i <= n; i++ {
		dp[i] = make([]int, amount+1)
	}

	for i := 0; i <= n; i++ {
		dp[i][0] = 1
	}

	for i := 0; i <= amount; i++ {
		dp[0][i] = math.MaxInt - 1
	}

	for i := 1; i <= amount; i++ {
		if i%coins[0] == 0 {
			dp[1][i] = 1
		} else {
			dp[1][i] = 0
		}
	}

	for i := 2; i <= n; i++ {
		for j := 1; j <= amount; j++ {
			if coins[i-1] <= j {
				dp[i][j] = dp[i][j-coins[i-1]] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	if dp[n][amount] == math.MaxInt-1 {
		fmt.Println(0)
	} else {
		fmt.Println(dp[n][amount])
	}
}
