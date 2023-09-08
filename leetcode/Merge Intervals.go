package leetcode

import "sort"

func merge_56(intervals [][]int) [][]int {
	n := len(intervals)
	sort.Slice(intervals, func(a, b int) bool {
		return intervals[a][0] < intervals[b][0]
	})

	res := make([][]int, 0, n)
	cur := intervals[0]
	for _, interval := range intervals[1:] {
		if interval[0] <= cur[1] {
			cur[1] = max(cur[1], interval[1])
		} else {
			res = append(res, cur)
			cur = interval
		}
	}
	res = append(res, cur)
	return res
}
