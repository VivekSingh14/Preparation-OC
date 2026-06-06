package main

import "fmt"

// longest repeating subsequence
// given string: AABEBCDD
// result = 3 or ABD

//solution approach

// we will find out longest common subsequence from same string with its copy
// but adding one more condition when i != j // line no: 41
func main9() {

	s1 := "AABEBCDD"
	//made same string
	s2 := "AABEBCDD"

	m := len(s1)
	n := len(s2)

	res := solveDP7(s1, s2, m, n)

	fmt.Println(res)
}

func solveDP7(s1 string, s2 string, m int, n int) int {

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
			if s1[i-1] == s2[j-1] && i != j {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[m][n]

}
