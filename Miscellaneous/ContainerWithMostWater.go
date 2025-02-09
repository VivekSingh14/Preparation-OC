package main

import "fmt"

func main5() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}

	fmt.Println(maxArea(height))
}

func maxArea(height []int) int {
	l := 0

	r := len(height) - 1
	max1 := 0
	for l < r {
		length := min1(height[l], height[r])
		if (r-l)*length > max1 && length == height[l] {
			max1 = (r - l) * length
			l++
		} else if (r-l)*length > max1 && length == height[r] {
			max1 = (r - l) * length
			r--
		} else {
			if length == height[l] {
				l++
			} else {
				r--
			}
		}
	}
	return max1
}

func min1(a, b int) int {
	if a > b {
		return b
	}
	return a
}
