package main

import (
	"fmt"
	"math"
)

var dp4 [7][12]int

// similar to count of subset sum from 01 knapsack
func main() {
	coins := []int{1, 2, 5}
	amount := 11

	for i := 0; i < len(coins)+1; i++ {
		for j := 0; j < amount+1; j++ {
			if i == 0 {
				dp4[i][j] = math.MaxInt
			}
			if j == 0 {
				dp4[i][j] = 0
			}
		}
	}

	res := knapsackDP3(coins, amount, len(coins))

	fmt.Println(res)

	//fmt.Println(dp4)

}

func knapsackDP3(coins []int, sum int, n int) int {

	for i := 1; i < n+1; i++ {
		for j := 1; j < sum+1; j++ {
			if coins[i-1] <= j {
				dp4[i][j] = min(1+dp4[i][j-coins[i-1]], dp4[i-1][j])
			} else {
				dp4[i][j] = dp4[i-1][j]
			}
		}
	}
	return dp4[n][sum]
}
