package main

import "fmt"

func main8() {
	text1 := "abcde"
	text2 := "ace"

	n := len(text1)
	m := len(text2)

	res := lcs(text1, text2, n, m)

	fmt.Println(res)

}

func lcs(text1 string, text2 string, n int, m int) int {

	if n == 0 || m == 0 {
		return 0
	}

	if text1[n-1] == text2[m-1] {
		return 1 + lcs(text1, text2, n-1, m-1)
	} else {
		return max(lcs(text1, text2, n-1, m), lcs(text1, text2, n, m-1))
	}
}
