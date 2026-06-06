package main

import "fmt"

func main8() {
	s1 := "acbcf"
	s2 := "abcdaf"

	//o/p := acbcdaf

	m := len(s1)
	n := len(s2)

	solveDP6(s1, s2, m, n)

}

func solveDP6(s1 string, s2 string, m int, n int) {

	var dp [7][8]int

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

	i := m
	j := n

	var resstr string
	for i > 0 && j > 0 {
		if s1[i-1] == s2[j-1] {
			resstr += string(s1[i-1])
			i--
			j--
		} else if dp[i][j-1] > dp[i-1][j] {
			resstr += string(s2[j-1])
			j--
		} else {
			resstr += string(s1[i-1])
			i--
		}
	}

	for i > 0 {
		resstr += string(s1[i-1])
		i--
	}

	for j > 0 {
		resstr += string(s2[j-1])
		j--
	}

	resstr = reverse12(resstr)

	fmt.Println(resstr)

}

func reverse12(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
