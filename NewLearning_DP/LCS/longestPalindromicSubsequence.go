package main

import "fmt"

func main6() {
	//single string is given, we neee to find out longest palindormic subsequence:  abcba
	s1 := "agbcba"

	// we will reverse s1 and assing it to s2 and find out the longest palindromic subsequence. reverse(s1)
	s2 := ReverseString(s1)

	m := len(s1)
	n := len(s2)

	res := solveDP4(s1, s2, m, n)

	fmt.Println(res)

}

func solveDP4(s1 string, s2 string, m int, n int) int {
	var dp [10][10]int

	for i := 0; i < m+1; i++ {
		for j := 0; j < n+1; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 0
			}
		}
	}

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]

}

func ReverseString(s string) string {
	runes := []rune(s)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	return string(runes)
}
