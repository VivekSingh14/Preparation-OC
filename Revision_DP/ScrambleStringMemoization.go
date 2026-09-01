package main

import (
	"fmt"
	"sort"
)

// declare
var map1 map[string]bool

func main19() {
	str1 := "abcde"
	str2 := "aecbd"

	fmt.Println(isScramble(str1, str2))
}

func isScramble(str1 string, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	//initialize the same map declared globally
	map1 = make(map[string]bool)

	return checkScramble(str1, str2)
}

func checkScramble(str1 string, str2 string) bool {

	if str1 == str2 {
		return true
	}
	key := str1 + "-" + str2

	if res, exists := map1[key]; exists {
		return res
	}

	x := []byte(str1)
	y := []byte(str2)

	sort.Slice(x, func(i, j int) bool {
		return x[i] < x[j]
	})

	sort.Slice(y, func(i, j int) bool {
		return y[i] < y[j]
	})

	if string(x) != string(y) {
		map1[key] = false
		return false
	}

	n := len(str1)

	for i := 1; i < n; i++ {
		noSwap := checkScramble(str1[:i], str2[:i]) && checkScramble(str1[i:], str2[i:])

		if noSwap {
			map1[key] = true
			return true
		}

		swap := checkScramble(str1[:i], str2[n-i:]) && checkScramble(str1[i:], str2[:n-i])

		if swap {
			map1[key] = true
			return true
		}
	}

	map1[key] = false

	return false

}
