package leetcode

import "sort"

func minProcessingTime(processorTime []int, tasks []int) int {
	res := 0
	sort.Ints(processorTime)
	sort.Slice(tasks, func(a, b int) bool {
		return tasks[a] > tasks[b]
	})
	for i, p := range processorTime {
		res = max(res, p+tasks[4*i])
	}
	return res
}
