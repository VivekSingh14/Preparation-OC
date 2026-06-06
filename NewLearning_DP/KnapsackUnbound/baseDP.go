package main

import "fmt"

var dp [5][5]int

func main2() {

	W := 4
	val := []int{1, 2, 3}
	wt := []int{4, 5, 1}

	for i := 0; i < len(val)+1; i++ {
		for j := 0; j < W+1; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 0
			}
		}
	}

	res := knapsackDP1(wt, val, W, len(val))

	fmt.Println(res)

}

func knapsackDP1(wt []int, val []int, W int, n int) int {

	for i := 1; i < n+1; i++ {
		for j := 1; j < W+1; j++ {
			if wt[i-1] <= j {
				//only dp[i-1][j-wt[i-1]] -> dp[i][j-wt[i-1]] for unbound knapsack
				dp[i][j] = max((val[i-1] + dp[i][j-wt[i-1]]), dp[i-1][j])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[n][W]
}
