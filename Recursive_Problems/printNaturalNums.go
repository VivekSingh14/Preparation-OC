package main

import "fmt"

func main() {
	n := 6

	print(n)
	fmt.Println("\n--------------")
	printreverse(n)
}

// increasing order
func print(n int) {
	if n == 1 {
		fmt.Print(n, " ")
		return
	}

	print(n - 1)

	fmt.Print(n, " ")
}

func printreverse(n int) {
	if n == 1 {
		fmt.Print(n, " ")
		return
	}

	fmt.Print(n, " ")

	printreverse(n - 1)
}
