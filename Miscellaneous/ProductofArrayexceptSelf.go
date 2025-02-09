package main

import "fmt"

func main1() {
	nums := []int{1, 2, 3, 4}

	res := []int{1, 1, 1, 1}

	for i := 1; i < len(res); i++ {
		res[i] = res[i-1] * nums[i-1]
	}
	rightProduct := 1
	for i := len(res) - 2; i >= 0; i-- {
		rightProduct *= nums[i+1]
		res[i] *= rightProduct
	}

	fmt.Println(res)
}
