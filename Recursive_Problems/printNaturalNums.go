package main

import "fmt"

func main() {
	n := 6

	print(n)
	fmt.Println("--------------")
	printreverse(n)
}

// increasing order
func print(n int) {
	if n == 0 {
		return
	}

	print(n - 1)

	fmt.Println(n)
}

func printreverse(n int) {
	if n == 0 {
		return
	}

	fmt.Println(n)

	printreverse(n - 1)
}
