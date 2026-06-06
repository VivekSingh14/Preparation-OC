package main

import "fmt"

func main3() {
	n := 5

	fmt.Println(NthFibonacci(n))

	dp := make([]int, n+1)

	for i := range n + 1 {
		dp[i] = -1
	}

	fmt.Println(NthFibonacciMemoiz(n, dp))

	fmt.Println(NthFibonacciIterative(n))

}

// recursion
func NthFibonacci(n int) int {

	if n <= 1 {
		return n
	}

	return NthFibonacci(n-1) + NthFibonacci(n-2)
}

// memoization
func NthFibonacciMemoiz(n int, dp []int) int {

	if n <= 1 {
		return n
	}

	if dp[n] != -1 {
		return dp[n]
	}

	dp[n] = NthFibonacciMemoiz(n-1, dp) + NthFibonacciMemoiz(n-2, dp)

	return dp[n]
}

func NthFibonacciIterative(n int) int {
	if n <= 1 {
		return n
	}

	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}
