package main

import (
	"fmt"
	"sort"
)

func main2() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}

	fmt.Println(merge(intervals))

}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := [][]int{intervals[0]}

	for k := 1; k < len(intervals); k++ {
		lastEnd := res[len(res)-1][1]
		if intervals[k][0] > lastEnd {
			res = append(res, intervals[k])
		} else {
			res[len(res)-1][1] = max(res[len(res)-1][1], intervals[k][1])
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
