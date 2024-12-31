// Happy Number
// Ex:  19 = 1+9^2 = 82 = 8^2+2^2 = 68 = 6^2+8^2 = 100 = 1^2 + 0^2 + 0^2 = 1[Happy Number]
package main

import "fmt"

func main() {
	num := 19

	fmt.Println(isHappy(num))
}

// Slow and fast pointer algo
func isHappy(n int) bool {
	slow := n
	fast := getNext(n)

	for fast != 1 && slow != fast {
		slow = getNext(slow)
		fast = getNext(getNext(fast))
	}
	return fast == 1
}

func getNext(n int) int {
	totatsum := 0

	for n > 0 {
		digit := n % 10
		n = n / 10
		totatsum += digit * digit
	}
	return totatsum
}
