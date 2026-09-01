package main

import "fmt"

var DP [201][201]int

func main1() {
	nums := []int{1, 5, 11, 5}

	sum := 0

	for i := range DP {
		for j := range DP[i] {
			DP[i][j] = -1
		}
	}

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	sum = sum / 2

	fmt.Println(CheckIfPartitionEqSum(nums, sum, len(nums)))

}

func CheckIfPartitionEqSum(nums []int, sum int, n int) bool {

	if sum == 0 {
		return true
	}

	if n == 0 {
		return false
	}

	if DP[sum][n-1] != -1 {
		return DP[sum][n-1] == 1
	}
	var temp bool
	if nums[n-1] <= sum {
		temp = CheckIfPartitionEqSum(nums, sum-nums[n-1], n-1) || CheckIfPartitionEqSum(nums, sum, n-1)

	} else {
		temp = CheckIfPartitionEqSum(nums, sum, n-1)
	}

	if temp {
		DP[sum][n] = 1
	} else {
		DP[sum][n] = 0
	}

	return temp
}
