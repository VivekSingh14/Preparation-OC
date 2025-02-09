package main

import "fmt"

func main2() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDisappearedNumbers(nums))
}

func findDisappearedNumbers(nums []int) []int {

	length := len(nums)

	if length == 0 {
		return []int{}
	}

	res := make([]int, length)

	for _, v := range nums {
		res[v-1] = v
	}
	i := 0
	for j := 0; j < len(res); j++ {
		if res[j] == 0 {
			res[i] = j + 1
			i++
		}
	}

	return res[:i]
}
