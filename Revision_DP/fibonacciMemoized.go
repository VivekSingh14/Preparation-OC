package main

import "fmt"

var dpfib []int

func main25() {

	num := 5

	dpfib = make([]int, num+1)

	dpfib[0] = 0
	dpfib[1] = 1

	fmt.Println(fibMem(num))
}

func fibMem(n int) int {
	if n <= 1 {
		return n
	}

	if dpfib[n] > 0 {
		return dpfib[n]
	}

	dpfib[n] = fibMem(n-1) + fibMem(n-2)

	return dpfib[n]
}
