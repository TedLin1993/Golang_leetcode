package leetcode

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := 0
	rightVal := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < rightVal {
			if intervals[i][1] < rightVal {
				rightVal = intervals[i][1]
			}
			res++
		} else {
			rightVal = intervals[i][1]
		}
	}
	return res
}
