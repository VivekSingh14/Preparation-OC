package main

import "fmt"

func main7() {
	//single string is given, we neee to find out longest palindormic subsequence:  abcba
	s1 := "agbcba"

	// we will reverse s1 and assing it to s2 and find out the longest common subsequence. reverse(s1)
	s2 := reverseString1(s1)

	m := len(s1)
	n := len(s2)

	res := solveDP5(s1, s2, m, n)

	//here to findout the number of deletion we get the difference between string length and lcs

	fmt.Println("minimum deletion to make the stirng palindrome is: ", m-res)

}

func solveDP5(s1 string, s2 string, m int, n int) int {
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

func reverseString1(s string) string {
	runes := []rune(s)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	return string(runes)
}
