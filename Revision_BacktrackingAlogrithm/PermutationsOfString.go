package main

import "fmt"

func main1() {

	str := "ABC"

	temp := []byte(str)

	start := 0

	permutate(temp, start)
}

func permutate(str []byte, start int) {
	if len(str) == start {
		fmt.Println(string(str))
		return
	}

	for i := start; i < len(str); i++ {
		//swap
		str[start], str[i] = str[i], str[start]
		//recursion
		permutate(str, start+1)
		//swap back for Backtracking
		str[start], str[i] = str[i], str[start]
	}
}
