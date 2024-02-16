package leetcode

import "container/heap"

func largestPerimeter_2971(nums []int) int64 {
	maxHeap := MaxHeap(nums)
	sum := 0
	for _, v := range nums {
		sum += v
	}
	heap.Init(&maxHeap)
	for len(maxHeap) > 0 {
		v := heap.Pop(&maxHeap).(int)
		sum -= v
		if sum > v {
			return int64(sum + v)
		}
	}
	return int64(-1)
}
