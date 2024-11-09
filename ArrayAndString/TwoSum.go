package main

import "fmt"

func main() {
	arr := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(arr, target))
}

func twoSum(nums []int, target int) []int {
	var res []int
	hashmap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if value, ok := hashmap[target-nums[i]]; ok {
			res = append(res, value)
			res = append(res, i)
		}
		hashmap[nums[i]] = i
	}
	return res
}
