package main

import "fmt"

func main3() {

	coins := []int{1, 2, 5}
	amount := 11

	fmt.Println(coinChange(coins, amount))

}

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = amount + 1
	}

	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] = min(dp[i], 1+dp[i-coin])
		}
	}

	if dp[amount] != amount+1 {
		return dp[amount]
	} else {
		return -1
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
