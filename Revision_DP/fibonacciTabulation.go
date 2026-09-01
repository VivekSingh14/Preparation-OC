package main

import "fmt"

var dpfib1 []int

func main26() {

	num := 5

	dpfib1 = make([]int, num+1)

	dpfib1[0] = 0
	dpfib1[1] = 1

	fmt.Println(fibTab(num))
}

func fibTab(n int) int {
	if n <= 1 {
		return n
	}

	if dpfib1[n] > 0 {
		return dpfib1[n]
	}

	for i := 2; i <= n; i++ {
		dpfib1[i] = dpfib1[i-1] + dpfib1[i-2]
	}

	return dpfib1[n]
}
