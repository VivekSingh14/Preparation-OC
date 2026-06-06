package main

import "fmt"

var dp1 [10][10]int

func main1() {

	N := 8
	price := []int{3, 5, 8, 9, 10, 17, 17, 20}
	//this is not given in problem
	rodLength := []int{1, 2, 3, 4, 5, 6, 7, 8}

	for i := 0; i < len(price)+1; i++ {
		for j := 0; j < N+1; j++ {
			if i == 0 || j == 0 {
				dp1[i][j] = 0
			}
		}
	}

	res := knapsackDP(rodLength, price, N, len(price))

	fmt.Println(res)

}

func knapsackDP(rodLength []int, price []int, N int, n int) int {

	for i := 1; i < n+1; i++ {
		for j := 1; j < N+1; j++ {
			if rodLength[i-1] <= j {
				//only dp1[i-1][j-wt[i-1]] -> dp1[i][j-wt[i-1]] for unbound knapsack
				dp1[i][j] = max((price[i-1] + dp1[i][j-rodLength[i-1]]), dp1[i-1][j])
			} else {
				dp1[i][j] = dp1[i-1][j]
			}
		}
	}
	return dp1[n][N]
}
