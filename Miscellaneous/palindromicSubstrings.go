package main

import "fmt"

func main11() {
	s := "abc"

	fmt.Println(countSubstrings(s))
}

func countSubstrings(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	count := len(s)

	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if isPalindrom(s, i, j) {
				count++
			}
		}
	}

	return count

}

func isPalindrom(s string, st, end int) bool {
	for st <= end {
		if s[st] != s[end] {
			return false
		}
		st++
		end--
	}
	return true
}
