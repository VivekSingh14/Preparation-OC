package main

import (
	"fmt"
	"sort"
)

func main4() {
	nums := []int{-1, 0, 1, 2, -1, -4}

	fmt.Println(threeSum(nums))

}

func threeSum(nums []int) [][]int {
	var res [][]int
	// sort nums
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		// skip duplicate
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// binary search
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum < 0 {
				l++
			} else if sum > 0 {
				r--
			} else {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			}
		}
	}
	return res
}
