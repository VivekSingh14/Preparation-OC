package main

import "fmt"

func main5() {

	nums := []int{100, 4, 200, 1, 3, 2}

	fmt.Println(longestConsecutive(nums))

}

func longestConsecutive(nums []int) int {
	set := map[int]bool{}
	for _, num := range nums {
		set[num] = true
	}

	res := 0
	for _, num := range nums {
		if set[num-1] {
			continue
		}

		sequence, temp := 1, num+1
		for set[temp] {
			sequence++
			temp++
		}

		res = max(res, sequence)
		if sequence > len(nums)/2 {
			break
		}
	}
	return res
}
func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}
