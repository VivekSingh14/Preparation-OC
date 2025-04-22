package main

import "fmt"

func main() {
	s := "aaa"

	fmt.Println(countSubstrings(s))
}

func countSubstrings(s string) int {

	res := 0

	i := 0

	for i < len(s) {
		l, r := i, i

		//odd length
		for l >= 0 && r < len(s) && s[l] == s[r] {
			res++
			l--
			r++
		}

		//even length

		l, r = i, i+1

		for l >= 0 && r < len(s) && s[l] == s[r] {
			res++
			l--
			r++
		}

		i++

	}

	return res
}
