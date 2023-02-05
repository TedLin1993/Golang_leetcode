package leetcode

import (
	"container/heap"
	"fmt"
	"math"
)

func pickGifts(gifts []int, k int) int64 {
	giftHeap := MaxHeap(gifts)
	heap.Init(&giftHeap)
	for i := 0; i < k; i++ {
		top := heap.Pop(&giftHeap).(int)
		top = int(math.Sqrt(float64(top)))
		heap.Push(&giftHeap, top)
	}
	res := 0
	for i := 0; i < len(giftHeap); i++ {
		res += giftHeap[i]
	}
	return int64(res)
}

func Test_pickGifts() {
	fmt.Println(pickGifts([]int{25, 64, 9, 4, 100}, 4)) //29
}
