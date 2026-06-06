package main

import "fmt"

var dp3 [7][11]int

// similar to count of subset sum from 01 knapsack
func main3() {
	coins := []int{2}
	amount := 3

	for i := 0; i < len(coins)+1; i++ {
		for j := 0; j < amount+1; j++ {
			if i == 0 {
				dp3[i][j] = 0
			}
			if j == 0 {
				dp3[i][j] = 1
			}
		}
	}

	res := knapsackDP2(coins, amount, len(coins))

	fmt.Println(res)

	fmt.Println(dp3)

}

func knapsackDP2(coins []int, sum int, n int) int {

	for i := 1; i < n+1; i++ {
		for j := 1; j < sum+1; j++ {
			if coins[i-1] <= j {
				dp3[i][j] = dp3[i][j-coins[i-1]] + dp3[i-1][j]
			} else {
				dp3[i][j] = dp3[i-1][j]
			}
		}
	}
	return dp3[n][sum]
}
