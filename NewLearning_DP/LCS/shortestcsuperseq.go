package main

import "fmt"

// nothing but, subtract the length of longest common subsequence from total length of both strings
func main4() {
	s1 := "geek"
	s2 := "eke"

	m := len(s1)
	n := len(s2)

	res := solveDP2(s1, s2, m, n)

	totallength := m + n

	fmt.Println("length of shortest common supersequence is: ", totallength-res)

}

func solveDP2(s1 string, s2 string, m int, n int) int {

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
