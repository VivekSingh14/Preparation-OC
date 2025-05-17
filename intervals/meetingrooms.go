package main

import (
	"fmt"
	"sort"
)

func main() {

	//true case
	intervals := [][]int{{0, 30}, {35, 40}, {45, 50}}

	//false case
	//intervals := [][]int{{0, 30}, {5, 10}, {15, 20}}

	fmt.Println(canAttendMeetings(intervals))

}

func canAttendMeetings(intervals [][]int) bool {

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	for i := 1; i < len(intervals); i++ {
		i1 := intervals[i-1]
		i2 := intervals[i]

		if i1[1] > i2[0] {
			return false
		}
	}
	return true
}
