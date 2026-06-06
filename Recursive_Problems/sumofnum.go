package main

import "fmt"

func main1() {

	n := 5

	fmt.Println(sum(n))
}

func sum(n int) int {
	if n == 0 {
		return 0
	}

	return n + sum(n-1)
}
