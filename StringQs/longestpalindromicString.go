package main

import "fmt"

func main2() {
	s := "babad"

	fmt.Println(longestPalindrome(s))

}

func longestPalindrome(s string) string {

	res := ""
	reslen := 0

	i := 0

	for i < len(s) {
		l, r := i, i

		//odd length
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > reslen {
				res = s[l : r+1]
				reslen = r - l + 1
			}
			l--
			r++
		}

		//even length

		l, r = i, i+1

		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > reslen {
				res = s[l : r+1]
				reslen = r - l + 1
			}
			l--
			r++
		}

		i++

	}

	return res
}
