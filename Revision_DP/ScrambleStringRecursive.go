package main

import (
	"fmt"
	"sort"
)

func main20() {
	str1 := "great"
	str2 := "rgeat"

	fmt.Println(isScramble2(str1, str2))
}

func isScramble2(str1 string, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	return checkScramble1(str1, str2)
}

func checkScramble1(str1 string, str2 string) bool {

	if str1 == str2 {
		return true
	}

	x := []byte(str1)
	y := []byte(str2)

	sort.Slice(x, func(i, j int) bool {
		return x[i] < x[j]
	})

	sort.Slice(y, func(i, j int) bool {
		return y[i] < y[j]
	})

	n := len(str1)

	for i := 1; i < n; i++ {
		noSwap := checkScramble1(str1[:i], str2[:i]) && checkScramble1(str1[i:], str2[i:])

		if noSwap {
			return true
		}

		swap := checkScramble1(str1[:i], str2[n-i:]) && checkScramble1(str1[i:], str2[:n-i])

		if swap {
			return true
		}
	}

	return false
}
