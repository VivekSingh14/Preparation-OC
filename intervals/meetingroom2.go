package main

import (
	"fmt"
	"sort"
)

func main() {

	intervals := [][]int{{0, 35}, {5, 10}, {15, 35}, {30, 40}}

	fmt.Println(minMeetingRooms(intervals))

}

func minMeetingRooms(intervals [][]int) int {
	var start []int

	for i := 0; i < len(intervals); i++ {
		start = append(start, intervals[i][0])
	}

	sort.Slice(start, func(i, j int) bool {
		return start[i] < start[j]
	})

	var end []int

	for i := 0; i < len(intervals); i++ {
		end = append(end, intervals[i][1])
	}

	sort.Slice(end, func(i, j int) bool {
		return end[i] < end[j]
	})

	res, count := 0, 0

	s, e := 0, 0

	for s < len(intervals) {
		if start[s] < end[e] {
			s++
			count++
		} else {
			e++
			count--
		}
		res = max1(res, count)
	}
	return res
}

func max1(a, b int) int {
	if a > b {
		return a
	}

	return b
}
