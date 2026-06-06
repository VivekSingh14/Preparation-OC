package main

import "fmt"

var t [1001][1001]int

func main1() {
	s1 := "AGGTAB"
	s2 := "GXTXAYB"

	for i := range 1001 {
		for j := range 1001 {
			t[i][j] = -1
		}
	}

	//res := lcsRecursive(s1, s2)

	lcsMemoize(s1, s2)

	//fmt.Println(t[len(s1)][len(s2)])

}

// func lcsRecursive(s1 string, s2 string) int {
// 	m := len(s1)
// 	n := len(s2)

// 	return solve(s1, s2, m, n)
// }

// func solve(s1 string, s2 string, m int, n int) int {

// 	if m == 0 || n == 0 {
// 		return 0
// 	}

// 	if s1[m-1] == s2[n-1] {
// 		return 1 + solve(s1, s2, m-1, n-1)
// 	}

// 	if s1[m-1] != s2[n-1] {
// 		return max(solve(s1, s2, m-1, n), solve(s1, s2, m, n-1))
// 	}
// 	return 0
// }

func lcsMemoize(s1 string, s2 string) {
	m := len(s1)
	n := len(s2)

	//solveMemoize(s1, s2, m, n)

	res := solveDP(s1, s2, m, n)

	fmt.Println(res)
}

func solveMemoize(s1 string, s2 string, m int, n int) int {

	if m == 0 || n == 0 {
		return 0
	}

	if t[m][n] != -1 {
		return t[m][n]
	}

	if s1[m-1] == s2[n-1] {
		t[m][n] = 1 + solveMemoize(s1, s2, m-1, n-1)
		return t[m][n]
	}

	if s1[m-1] != s2[n-1] {

		t[m][n] = max(solveMemoize(s1, s2, m-1, n), solveMemoize(s1, s2, m, n-1))
		return t[m][n]
	}

	return 0
}

func solveDP(s1 string, s2 string, m int, n int) int {

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
