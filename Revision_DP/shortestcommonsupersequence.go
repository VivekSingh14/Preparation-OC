package main

import "fmt"

// first we will findout longest common subsequence.
// total len of shortest common supersequence would be m+n-len(lcs)
// but here we have to print the scs.

func main11() {

	dp2 := make([][]int, 1001)

	for x := range dp2 {
		dp2[x] = make([]int, 1001)
	}

	str1 := "abac"
	str2 := "cab"

	m := len(str1)
	n := len(str2)

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if str1[i-1] == str2[j-1] {
				dp2[i][j] = 1 + dp2[i-1][j-1]
			} else {
				dp2[i][j] = max(dp2[i-1][j], dp2[i][j-1])
			}
		}
	}

	var res string

	i := m
	j := n

	for i > 0 && j > 0 {
		if str1[i-1] == str2[j-1] {
			res += string(str1[i-1])
			i--
			j--
		} else if dp2[i][j-1] > dp2[i-1][j] {
			res += string(str2[j-1])
			j--
		} else {
			res += string(str1[i-1])
			i--
		}
	}

	for i > 0 {
		res += string(str1[i-1])
		i--
	}

	for j > 0 {
		res += string(str2[j-1])
		j--
	}

	res = Reverse(res)

	fmt.Println(res)

}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
