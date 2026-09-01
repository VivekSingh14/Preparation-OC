package main

import "fmt"

var dp [1001][1001]int

func main9() {

	text1 := "abcde"
	text2 := "ace"

	for x := range dp {
		for y := range dp[x] {
			dp[x][y] = -1
		}
	}

	m := len(text1)
	n := len(text2)

	res := lcsMemoized(text1, text2, m, n)

	fmt.Println(res)

}

func lcsMemoized(text1 string, text2 string, m int, n int) int {

	if m == 0 || n == 0 {
		return 0
	}

	if text1[m-1] == text2[n-1] {
		dp[m][n] = 1 + lcsMemoized(text1, text2, m-1, n-1)
	} else {
		dp[m][n] = max(lcsMemoized(text1, text2, m-1, n), lcsMemoized(text1, text2, m, n-1))
	}

	return dp[m][n]
}
