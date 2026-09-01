package main

import "fmt"

func main2() {
	nums := []int{1, 2, 5, 2}

	sum := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	if sum%2 != 0 {
		fmt.Println(sum == 0)
		return
	}

	sum /= 2

	dp2 := make([][]bool, len(nums)+1)

	for i := 0; i < len(nums)+1; i++ {
		dp2[i] = make([]bool, sum+1)
	}

	// --------initializing when sum=0 and nums are available, so always empty set is available-----------
	for i := 0; i < len(dp2); i++ {
		for j := 0; j < len(dp2[i]); j++ {
			if j == 0 {
				dp2[i][j] = true
			}
		}
	}

	// for i := 0; i < len(dp2); i++ {
	// 	for j := 0; j < len(dp2[i]); j++ {
	// 		fmt.Print(dp2[i][j], " ")
	// 	}
	// 	fmt.Println()
	// }

	fmt.Println(ifEqualSumsubsetPresent(dp2, nums))

	for i := 0; i < len(dp2); i++ {
		for j := 0; j < len(dp2[i]); j++ {
			fmt.Print(dp2[i][j], " ")
		}
		fmt.Println()
	}

}

func ifEqualSumsubsetPresent(dp [][]bool, nums []int) bool {

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if nums[i-1] <= j {
				dp[i][j] = dp[i-1][j-nums[i-1]] || dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[len(dp)-1][len(dp[0])-1]
}
