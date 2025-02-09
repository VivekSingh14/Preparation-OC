package main

import "fmt"

func main5() {
	text1 := "abcde"
	text2 := "ace"

	fmt.Println(longestCommonSubsequence(text1, text2))
}

func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := range m + 1 {
		dp[i] = make([]int, n+1)
	}
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if text1[i] == text2[j] {
				dp[i][j] = 1 + dp[i+1][j+1]
			} else {
				dp[i][j] = Max(dp[i][j+1], dp[i+1][j])
			}
		}
	}
	return dp[0][0]
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
