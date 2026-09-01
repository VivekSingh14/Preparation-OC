package main

import "fmt"

func main24() {

	num := 5

	fmt.Println(fib(num))
}

func fib(n int) int {
	if n <= 1 {
		return n
	}

	return fib(n-1) + fib(n-2)
}
