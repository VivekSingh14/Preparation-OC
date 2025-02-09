package main

import "fmt"

// House robber -> house in circular route.
func main() {

	nums := []int{1, 2, 3}

	fmt.Println(rob2(nums))

}

func rob2(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}

	max1 := rob1(nums[1:])

	max2 := rob1(nums[:len(nums)-1])

	if max1 > max2 {
		return max1
	} else {
		return max2
	}
}

func rob1(nums []int) int {
	rob1, rob2 := 0, 0

	for _, num := range nums {
		temp := max1(num+rob1, rob2)
		rob1 = rob2
		rob2 = temp
	}
	return rob2
}

func max1(a, b int) int {
	if a > b {
		return a
	}

	return b
}
