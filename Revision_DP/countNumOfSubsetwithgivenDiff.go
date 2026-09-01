package main

import "fmt"

// -------- Count the number of subset with given difference -----
// ---------                 OR                              -----
// ---------    diff SUM           -----------------------------
func main3() {

	arr := []int{1, 1, 2, 3} //nums as per leetcode
	diff := 1                //target

	sum := 0

	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}

	if abs(diff) > sum {
		fmt.Println("0")
	}

	if (sum+diff)%2 != 0 {
		fmt.Println("0")
	}

	diff = (sum + diff) / 2

	dp := make([][]int, len(arr)+1)

	for i := 0; i <= len(arr); i++ {
		dp[i] = make([]int, diff+1)
	}

	dp[0][0] = 1

	for i := 1; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			if arr[i-1] <= j {
				dp[i][j] = dp[i-1][j] + dp[i-1][j-arr[i-1]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			fmt.Print(dp[i][j], " ")
		}
		fmt.Println()
	}

	fmt.Println(dp[len(arr)][diff])

	fmt.Println(len(arr), " ", diff)

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
