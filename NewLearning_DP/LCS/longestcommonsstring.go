package main

import "fmt"

var dp [10][10]int

func main2() {

	s1 := "ABC"
	s2 := "ABC"

	l1 := len(s1)
	l2 := len(s2)

	for i := range 10 {
		for j := range 10 {
			dp[i][j] = -1
		}
	}

	//res := ans(s1, s2, l1, l2, 0)
	//fmt.Println(res)

	//res := ansMemoize(s1, s2, l1, l2, 0)
	res := ansDP(s1, s2, l1, l2)
	fmt.Println(res)

	//fmt.Println(dp)

}

func ans(s1 string, s2 string, n int, m int, leng int) int {

	if n == 0 || m == 0 {
		return 0
	}
	if s1[n-1] == s2[m-1] {
		leng = ans(s1, s2, n-1, m-1, leng+1)
	}
	return max(leng, ans(s1, s2, n-1, m, 0), ans(s1, s2, n, m-1, 0))
}

func ansMemoize(s1 string, s2 string, n int, m int, leng int) int {
	if n == 0 || m == 0 {
		return 0
	}

	if dp[m][n] != -1 {
		return dp[m][n]
	}

	if s1[m-1] == s2[n-1] {
		dp[m][n] = ansMemoize(s1, s2, m-1, n-1, leng+1)
		return dp[m][n]
	} else {
		dp[m][n] = max(leng, ans(s1, s2, n-1, m, 0), ans(s1, s2, n, m-1, 0))
		return dp[m][n]
	}
}

func ansDP(s1 string, s2 string, n int, m int) int {
	var dp1 [4][4]int

	for i := 0; i < m+1; i++ {
		for j := 0; j < n+1; j++ {
			if i == 0 || j == 0 {
				dp1[i][j] = 0
			}
		}
	}
	var short_max int
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if s1[i-1] == s2[j-1] {
				dp1[i][j] = 1 + dp1[i-1][j-1]
				short_max = max(short_max, dp1[i][j])
			} else {
				dp1[i][j] = 0
			}
		}
	}

	return short_max
}
