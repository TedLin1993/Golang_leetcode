package leetcode

import (
	"container/heap"
	"fmt"
)

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	n1, n2 := len(nums1), len(nums2)
	maxPairHeap := MaxPairHeap{}
	for i := 0; i < n1; i++ {
		for j := 0; j < n2; j++ {
			sum := nums1[i] + nums2[j]
			if len(maxPairHeap) < k {
				heap.Push(&maxPairHeap, []int{nums1[i], nums2[j]})
			} else if sum < maxPairHeap[0][0]+maxPairHeap[0][1] {
				maxPairHeap[0] = []int{nums1[i], nums2[j]}
				heap.Fix(&maxPairHeap, 0)
			} else {
				break
			}
		}
	}
	return maxPairHeap
}

type MaxPairHeap [][]int

func (h MaxPairHeap) Len() int           { return len(h) }
func (h MaxPairHeap) Less(i, j int) bool { return h[i][0]+h[i][1] > h[j][0]+h[j][1] }
func (h MaxPairHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxPairHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}
func (h *MaxPairHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func Test_kSmallestPairs() {
	fmt.Println(kSmallestPairs([]int{1, 2, 4, 5, 6}, []int{3, 5, 7, 9}, 3)) //[[1,3],[2,3],[1,5]]
}
