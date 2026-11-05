package main

import "fmt"

func main1() {

	arr := []int{2, 7, 11, 15}
	target := 9

	fmt.Println(twoSum(arr, target))
}

func twoSum(arr []int, target int) []int {
	first := 0

	last := len(arr) - 1

	res := make([]int, 0)

	for first < last {
		if arr[first]+arr[last] == target {
			res = append(res, first+1)
			res = append(res, last+1)
			return res
		} else if arr[first]+arr[last] > target {
			last--
		} else {
			first++
		}
	}
	return nil
}
