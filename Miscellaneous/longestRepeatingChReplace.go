package main

import "fmt"

func main6() {
	s := "AABABBA"
	k := 1

	fmt.Println(characterReplacement(s, k))
}

func characterReplacement(s string, k int) int {
	freq := map[byte]int{}
	maxFreq := 1
	maxLen := 0

	l := 0
	for r := l; r < len(s); r++ {

		freq[s[r]]++

		maxFreq = max(maxFreq, freq[s[r]])

		if (r-l+1)-maxFreq > k {
			freq[s[l]]--
			l++
		}

		maxLen = max(maxLen, r-l+1)
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
