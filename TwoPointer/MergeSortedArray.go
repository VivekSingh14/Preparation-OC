package main

import "fmt"

func main() {

	// nums1 := []int{1, 2, 3, 0, 0, 0}
	// m := 3
	// nums2 := []int{2, 5, 6}
	// n := 3

	nums1 := []int{0}
	m := 0
	nums2 := []int{1}
	n := 1

	merge(nums1, m, nums2, n)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	mergedNums := make([]int, m+n)
	i := 0
	j := 0
	k := 0

	for i < m && j < n {
		if nums1[i] < nums2[j] {
			mergedNums[k] = nums1[i]
			i++
		} else {
			mergedNums[k] = nums2[j]
			j++
		}
		k++
	}

	for i < m {
		mergedNums[k] = nums1[i]
		k++
		i++
	}

	for j < n {
		mergedNums[k] = nums2[j]
		k++
		j++
	}

	fmt.Println(mergedNums)
}
