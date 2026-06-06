package main

import "fmt"

// This is 0-1 Knapsack algo
func main1() {

	nums := []int{1, 4, 11, 5}

	capacity := 0

	for i := 0; i < len(nums); i++ {
		capacity += nums[i]
	}

	fmt.Println(canPartition(nums, capacity/2, len(nums)))

	fmt.Println(canPartitionMem(nums))
}

// Recursive approach
func canPartition(nums []int, capacity int, n int) bool {

	if capacity == 0 {
		return true
	}

	if n == 0 {
		return false
	}

	if nums[n-1] > capacity {
		return canPartition(nums, capacity, n-1)
	}

	return canPartition(nums, capacity, n-1) || canPartition(nums, capacity-nums[n-1], n-1)
}

// Memoization method
func canPartitionMem(nums []int) bool {

	capacity := 0

	for i := 0; i < len(nums); i++ {
		capacity += nums[i]
	}

	if capacity%2 != 0 {
		return false
	}

	capacity /= 2

	n := len(nums)

	dp, visited := make([][]bool, capacity+1), make([][]bool, capacity+1)
	for i := 0; i < capacity+1; i++ {
		dp[i] = make([]bool, n)
	}
	for i := 0; i < capacity+1; i++ {
		visited[i] = make([]bool, n)
	}

	return solve(nums, dp, visited, capacity, n-1)
}

func solve(nums []int, dp [][]bool, visited [][]bool, capacity int, n int) bool {
	if n < 0 || capacity < 0 {
		return false
	}

	if capacity == 0 {
		dp[capacity][n] = true
		return dp[capacity][n]
	}

	if visited[capacity][n] {
		return dp[capacity][n]
	}

	visited[capacity][n] = true
	dp[capacity][n] = solve(nums, dp, visited, capacity, n-1) || solve(nums, dp, visited, capacity-nums[n], n-1)
	return dp[capacity][n]
}
