package leetcode

import (
	"container/heap"
	"fmt"
)

func totalCost(costs []int, k int, candidates int) int64 {
	n := len(costs)
	left, right := candidates-1, max(candidates, n-candidates)
	leftHeap := MinHeap(costs[:left+1])
	rightHeap := MinHeap(costs[right:])
	heap.Init(&leftHeap)
	heap.Init(&rightHeap)

	res := 0
	for i := 0; i < k; i++ {
		if len(leftHeap) != 0 && (len(rightHeap) == 0 || leftHeap[0] <= rightHeap[0]) {
			res += leftHeap[0]
			heap.Pop(&leftHeap)
			if left < right-1 {
				left++
				heap.Push(&leftHeap, costs[left])
			}
		} else {
			res += rightHeap[0]
			heap.Pop(&rightHeap)
			if left < right-1 {
				right--
				heap.Push(&rightHeap, costs[right])
			}
		}
	}
	return int64(res)
}

func Test_totalCost() {
	fmt.Println(totalCost([]int{17, 12, 10, 2, 7, 2, 11, 20, 8}, 3, 4))                                 //11
	fmt.Println(totalCost([]int{1, 2, 4, 1}, 3, 3))                                                     //4
	fmt.Println(totalCost([]int{1, 2, 4, 1}, 3, 4))                                                     //4
	fmt.Println(totalCost([]int{31, 25, 72, 79, 74, 65, 84, 91, 18, 59, 27, 9, 81, 33, 17, 58}, 11, 2)) //423
	fmt.Println(totalCost([]int{57, 33, 26, 76, 14, 67, 24, 90, 72, 37, 30}, 11, 2))                    //526
}
