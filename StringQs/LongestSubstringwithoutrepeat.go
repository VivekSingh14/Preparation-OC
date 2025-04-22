package main

import "fmt"

func main1() {
	s := "pwwkew"

	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {

	loc_min := 0
	loc_max := 0
	max_window_size1 := 0
	//max_window_star := 0

	countChar := make([]int, 26)
	countChar[s[0]-'a']++

	for i := 1; i < len(s); i++ {
		countChar[s[i]-'a']++
		loc_max++
		for !isUnique(countChar) {
			countChar[s[loc_min]-'a']--
			loc_min++
		}
		if loc_max-loc_min+1 > max_window_size1 {
			max_window_size1 = loc_max - loc_min + 1
			//max_window_star = loc_min
		}
	}

	//fmt.Println(max_window_size1)
	//fmt.Println(max_window_star)
	//fmt.Println(s[max_window_star : max_window_size1+max_window_star])

	return max_window_size1
}

func isUnique(countchars []int) bool {
	for i := 0; i < len(countchars); i++ {
		if countchars[i] > 1 {
			return false
		}
	}
	return true
}
