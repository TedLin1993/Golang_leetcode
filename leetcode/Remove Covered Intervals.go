package leetcode

import (
	"fmt"
	"sort"
)

func removeCoveredIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] < intervals[j][0] {
			return true
		}
		if intervals[i][0] > intervals[j][0] {
			return false
		}
		return intervals[i][1] > intervals[j][1]
	})

	res := len(intervals)
	rightVal := 0
	for i := 0; i < len(intervals); i++ {
		if intervals[i][1] <= rightVal {
			res--
		} else {
			rightVal = intervals[i][1]
		}
	}

	return res
}

func TestRemoveCoveredIntervals() {
	fmt.Println(removeCoveredIntervals([][]int{{1, 4}, {3, 6}, {2, 8}})) //2
	fmt.Println(removeCoveredIntervals([][]int{{1, 4}, {2, 3}}))         //1
	fmt.Println(removeCoveredIntervals([][]int{{3, 4}, {1, 2}, {1, 4}})) //1
}
