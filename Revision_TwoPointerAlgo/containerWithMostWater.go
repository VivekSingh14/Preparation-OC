package main

import "fmt"

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}

	first := 0

	last := len(height) - 1

	max := 0

	for first < last {
		length := min(height[first], height[last])
		if (last-first)*length > max && length == height[first] {
			max = (last - first) * length
			first++
		} else if (last-first)*length > max && length == height[last] {
			max = (last - first) * length
			last--
		} else if length == height[first] {
			first++
		} else {
			last--
		}
	}
	fmt.Println(max)
}
