package leetcode

import (
	"container/heap"
	"fmt"
)

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func furthestBuilding(heights []int, bricks int, ladders int) int {
	largestGaps := MinHeap{}
	for i := 1; i < len(heights); i++ {
		gap := heights[i] - heights[i-1]
		if gap <= 0 {
			continue
		}

		if ladders != 0 {
			if largestGaps.Len() < ladders {
				heap.Push(&largestGaps, gap)
			} else if gap <= largestGaps[0] {
				bricks -= gap
			} else {
				bricks -= heap.Pop(&largestGaps).(int)
				heap.Push(&largestGaps, gap)
			}
		} else {
			bricks -= gap
		}

		if bricks < 0 {
			return i - 1
		}
	}

	return len(heights) - 1
}

func TestfurthestBuilding() {
	fmt.Println(furthestBuilding([]int{4, 2, 7, 6, 9, 14, 12}, 5, 1))          //4
	fmt.Println(furthestBuilding([]int{4, 12, 2, 7, 3, 18, 20, 3, 19}, 10, 2)) //7
	fmt.Println(furthestBuilding([]int{14, 3, 19, 3}, 17, 0))                  //3
	fmt.Println(furthestBuilding([]int{1, 2}, 0, 0))                           //0
}
