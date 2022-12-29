package leetcode

import (
	"container/heap"
	"fmt"
)

func minStoneSum(piles []int, k int) int {
	maxHeap := MaxHeap(piles)
	heap.Init(&maxHeap)
	for i := 0; i < k; i++ {
		maxHeap[0] = maxHeap[0] - maxHeap[0]>>1
		heap.Fix(&maxHeap, 0)
	}

	res := 0
	for _, pile := range maxHeap {
		res += pile
	}
	return res
}

func Test_minStoneSum() {
	fmt.Println(minStoneSum([]int{5, 4, 9}, 2))    //12
	fmt.Println(minStoneSum([]int{4, 3, 6, 7}, 3)) //12
}
