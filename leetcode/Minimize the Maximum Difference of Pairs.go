package leetcode

import (
	"container/heap"
	"fmt"
	"sort"
)

func minimizeMax(nums []int, p int) int {
	if p == 0 {
		return 0
	}
	sort.Ints(nums)
	maxHeap := MaxHeap{}
	for i := 0; i < len(nums)-1; i = i + 2 {
		diff := nums[i+1] - nums[i]
		if len(maxHeap) < p {
			heap.Push(&maxHeap, diff)
		} else {
			if diff < maxHeap[0] {
				heap.Pop(&maxHeap)
				heap.Push(&maxHeap, diff)
			}
		}
	}
	return maxHeap[0]
}

func Test_minimizeMax() {
	fmt.Println(minimizeMax([]int{4, 2, 1, 2}, 1))          //0
	fmt.Println(minimizeMax([]int{0, 5, 3, 4}, 0))          //0
	fmt.Println(minimizeMax([]int{3, 4, 2, 3, 2, 1, 2}, 3)) //1
}
