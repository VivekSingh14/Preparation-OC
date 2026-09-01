package main

import "fmt"

func main10() {

	text1 := "abcde"
	text2 := "ace"

	m := len(text1)
	n := len(text2)

	dp1 := make([][]int, m+1)

	for x := range dp1 {
		dp1[x] = make([]int, n+1)
	}

	//already initialized first and first column with default value 0

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp1[i][j] = 1 + dp1[i-1][j-1]
			} else {
				dp1[i][j] = max(dp1[i][j-1], dp1[i-1][j])
			}
		}
	}

	fmt.Println(dp1[m][n])
}
