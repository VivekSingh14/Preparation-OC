package main

import "fmt"

var dp3 [7][11]int

// given {2,3,5,6,8,10}  subset with given sums[10] are: {2,3,5}, {2,8} {10} -> 3
func main5() {
	nums := []int{2, 3, 5, 6, 8, 10}
	sum := 10

	for i := 0; i < len(nums)+1; i++ {
		for j := 0; j < sum+1; j++ {
			if i == 0 {
				dp3[i][j] = 0
			}
			if j == 0 {
				dp3[i][j] = 1
			}
		}
	}

	res := knapsackDP1(nums, sum, len(nums))

	fmt.Println(res)

	fmt.Println(dp3)

}

func knapsackDP1(nums []int, sum int, n int) int {

	for i := 1; i < n+1; i++ {
		for j := 1; j < sum+1; j++ {
			if nums[i-1] <= j {
				dp3[i][j] = dp3[i-1][j-nums[i-1]] + dp3[i-1][j]
			} else {
				dp3[i][j] = dp3[i-1][j]
			}
		}
	}
	return dp3[n][sum]
}
