package main

// import "fmt"

// // o/p := 3

// var dp [5][5]int
// var dp1 [5][5]int

// func main3() {
// 	W := 4
// 	val := []int{1, 2, 3}
// 	wt := []int{4, 5, 1}

// 	// res := knapsackrecursive(wt, val, W, len(val))

// 	// fmt.Println(res)

// 	// for i := 0; i < len(dp); i++ {
// 	// 	for j := 0; j < len(dp[i]); j++ {
// 	// 		dp[i][j] = -1
// 	// 	}
// 	// }

// 	//knapsackmemoize(wt, val, W, len(val))

// 	//fmt.Println(dp[len(val)][W])

// 	for i := 0; i < len(val)+1; i++ {
// 		for j := 0; j < W+1; j++ {
// 			if i == 0 || j == 0 {
// 				dp1[i][j] = 0
// 			}
// 		}
// 	}

// 	res := knapsackDP(wt, val, W, len(val))

// 	fmt.Println(res)

// }

// func knapsackrecursive(wt []int, val []int, W int, n int) int {

// 	if n == 0 || W == 0 {
// 		return 0
// 	}

// 	if wt[n-1] <= W {
// 		return max(val[n-1]+knapsackrecursive(wt, val, W-wt[n-1], n-1), knapsackrecursive(wt, val, W, n-1))
// 	} else if wt[n-1] > W {
// 		return knapsackrecursive(wt, val, W, n-1)
// 	}
// 	return 0
// }

// func knapsackmemoize(wt []int, val []int, W int, n int) int {

// 	if n == 0 || W == 0 {
// 		return 0
// 	}
// 	if dp[n][W] != -1 {
// 		return dp[n][W]
// 	}

// 	if wt[n-1] <= W {
// 		dp[n][W] = max(val[n-1]+knapsackmemoize(wt, val, W-wt[n-1], n-1), knapsackmemoize(wt, val, W, n-1))
// 		return dp[n][W]
// 	} else if wt[n-1] > W {
// 		dp[n][W] = knapsackmemoize(wt, val, W, n-1)
// 		return dp[n][W]

// 	}
// 	return 0
// }

// func knapsackDP(wt []int, val []int, W int, n int) int {

// 	for i := 1; i < n+1; i++ {
// 		for j := 1; j < W+1; j++ {
// 			if wt[i-1] <= j {
// 				dp1[i][j] = max((val[i-1] + dp1[i-1][j-wt[i-1]]), dp1[i-1][j])
// 			} else {
// 				dp1[i][j] = dp1[i-1][j]
// 			}
// 		}
// 	}
// 	return dp1[n][W]
// }
