package main

import "fmt"

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"

	fmt.Println(minWindow(s, t))

}

func minWindow(s string, t string) string {
	m, n := len(s), len(t)
	if m < n {
		return ""
	}

	// Frequency map to store the count of characters in string t
	freq := make(map[byte]int)

	// Count of unique characters in string t
	uniqueChars := 0

	// Initialize the frequency map and unique character count
	for i := 0; i < n; i++ {
		freq[t[i]]++
		if freq[t[i]] == 1 {
			uniqueChars++
		}
	}

	// Pointers i and j for sliding window
	// start to store the starting index of the minimum window
	// len for the length of the minimum window
	var i, j int
	start, len := -1, m

	// Move the right pointer j until all required characters are included
	for j < m {
		for j < m && uniqueChars > 0 {
			freq[s[j]]--
			if freq[s[j]] == 0 {
				uniqueChars--
			}

			j++
		}

		// Move the left pointer i to find the minimum window
		for uniqueChars == 0 {
			if j-i <= len {
				start = i
				len = j - i
			}

			freq[s[i]]++
			if freq[s[i]] == 1 {
				uniqueChars++
			}

			i++
		}
	}

	// Return the minimum window substring if found, otherwise an empty string
	if start != -1 {
		return s[start : start+len]
	}

	return ""
}
