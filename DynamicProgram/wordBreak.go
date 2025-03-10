package main

import "fmt"

func main6() {
	s := "leetcode"
	wordDict := []string{"neet", "code"}

	fmt.Println(wordBreak(s, wordDict))
}

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[len(s)] = true

	for i := len(s) - 1; i >= 0; i-- {
		for _, word := range wordDict {
			if (i+len(word) <= len(s)) && (s[i:i+len(word)] == word) {
				dp[i] = dp[i+len(word)]
			}
			if dp[i] {
				break
			}
		}
	}

	return dp[0]
}
