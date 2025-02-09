package main

import "fmt"

//Two Sum to target

func main1() {
	arr := []int{2, 7, 11, 15, 17, 19, 24, 28, 33}
	target := 18
	fmt.Println(twoSum(arr, target))

	fmt.Println(twoPointer(arr, target))
}

// First Approach
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

// Second Approach using two pointer algo
func twoPointer(nums []int, target int) []int {
	var res []int
	end := len(nums) - 1
	for start := 0; start < end; {
		if nums[start]+nums[end] < target {
			start++
		} else if nums[start]+nums[end] > target {
			end--
		} else {
			res = append(res, start)
			res = append(res, end)
			return res
		}
	}
	return nil
}
