package leetcode

import (
	"fmt"
	"sort"
)

func insert(intervals [][]int, newInterval []int) [][]int {
	idx := sort.Search(len(intervals), func(i int) bool { return newInterval[0] <= intervals[i][0] })
	if idx == len(intervals) {
		intervals = append(intervals, newInterval)
	} else {
		intervals = append(intervals[:idx+1], intervals[idx:]...)
		intervals[idx] = newInterval
	}
	res := make([][]int, 0, len(intervals)+1)
	res = append(res, intervals[0])
	for i := 1; i < len(intervals); i++ {
		interval := intervals[i]
		lastInterval := res[len(res)-1]
		if interval[0] > lastInterval[1] {
			res = append(res, interval)
		} else {
			lastInterval[1] = max(interval[1], lastInterval[1])
		}
	}

	return res
}

func Test_insert() {
	fmt.Println(insert([][]int{{1, 3}, {6, 9}}, []int{2, 5}))                            //[[1,5],[6,9]]
	fmt.Println(insert([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8})) //[[1,2],[3,10],[12,16]]
	fmt.Println(insert([][]int{{1, 3}, {6, 9}}, []int{4, 5}))                            //[[1,3],[4,5],[6,9]]
	fmt.Println(insert([][]int{{1, 3}, {6, 9}}, []int{0, 1}))                            //[[0 3] [6 9]]
	fmt.Println(insert([][]int{{1, 3}, {6, 9}}, []int{7, 10}))                           //[[1 3] [6 10]]
}
