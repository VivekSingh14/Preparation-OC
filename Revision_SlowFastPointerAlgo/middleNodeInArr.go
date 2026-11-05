package main

import "fmt"

func main1() {
	arr := []int{1, 2, 3, 4, 5}

	res := middleNode(arr)

	fmt.Println(res)
}

func middleNode(arr []int) int {

	end := len(arr) - 1

	if end <= 1 {
		return arr[0]
	}
	slow := 0
	fast := 0

	for fast < end {
		slow++
		fast = fast + 2
	}

	return arr[slow]
}
