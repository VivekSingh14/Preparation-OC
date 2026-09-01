package main

import "fmt"

func main15() {

	s := "cbbd"
	s2 := Reverse2(s)

	m := len(s)
	n := len(s2)

	dp4 := make([][]int, 1001)

	for x := range dp4 {
		dp4[x] = make([]int, 1001)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[i-1] == s2[j-1] {
				dp4[i][j] = 1 + dp4[i-1][j-1]
			} else {
				dp4[i][j] = max(dp4[i-1][j], dp4[i][j-1])
			}
		}
	}

	fmt.Println(dp4[m][n])
}

func Reverse2(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
