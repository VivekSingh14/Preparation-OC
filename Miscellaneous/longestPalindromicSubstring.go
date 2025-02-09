package main

import "fmt"

func main() {
	s := "babad"

	fmt.Println(longestPalindrome(s))
}

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	maxlen := 0
	start := 0
	end := 0
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if IsPalindrom(s, i, j) && maxlen < j-i+1 {
				maxlen = j - i + 1
				start = i
				end = j
			}
		}
	}
	res := []rune(s)
	res = res[start : end+1]
	return string(res)
}

func IsPalindrom(s string, st, end int) bool {
	for st <= end {
		if s[st] != s[end] {
			return false
		}
		st++
		end--
	}
	return true
}
