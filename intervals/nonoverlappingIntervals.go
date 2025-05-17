package main

import (
	"fmt"
	"sort"
)

func main3() {

	intervals := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}

	fmt.Println(eraseOverlapIntervals(intervals))
}

func eraseOverlapIntervals(intervals [][]int) int {
	//sorting based on the start value
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	res := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < intervals[i-1][1] {
			intervals[i] = intervals[i-1]
			res++
		}
	}

	return res
}
