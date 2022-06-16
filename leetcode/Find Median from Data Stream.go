package leetcode

// import (
// 	"container/heap"
// 	"fmt"
// )

// type MinHeap []int

// func (h MinHeap) Len() int           { return len(h) }
// func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
// func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// func (h *MinHeap) Push(x any) {
// 	// Push and Pop use pointer receivers because they modify the slice's length,
// 	// not just its contents.
// 	*h = append(*h, x.(int))
// }

// func (h *MinHeap) Pop() any {
// 	old := *h
// 	n := len(old)
// 	x := old[n-1]
// 	*h = old[0 : n-1]
// 	return x
// }

// type MaxHeap []int

// func (h MaxHeap) Len() int           { return len(h) }
// func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
// func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// func (h *MaxHeap) Push(x any) {
// 	// Push and Pop use pointer receivers because they modify the slice's length,
// 	// not just its contents.
// 	*h = append(*h, x.(int))
// }

// func (h *MaxHeap) Pop() any {
// 	old := *h
// 	n := len(old)
// 	x := old[n-1]
// 	*h = old[0 : n-1]
// 	return x
// }

// type MedianFinder struct {
// 	large MinHeap
// 	small MaxHeap
// }

// func Constructor() MedianFinder {
// 	return MedianFinder{}
// }

// func (this *MedianFinder) AddNum(num int) {
// 	if this.small.Len() == 0 {
// 		this.small.Push(num)
// 		return
// 	}

// 	if num <= this.small[0] {
// 		heap.Push(&this.small, num)
// 		if this.small.Len() > this.large.Len()+1 {
// 			v := heap.Pop(&this.small)
// 			heap.Push(&this.large, v)
// 		}
// 	} else {
// 		heap.Push(&this.large, num)
// 		if this.large.Len() > this.small.Len()+1 {
// 			v := heap.Pop(&this.large)
// 			heap.Push(&this.small, v)
// 		}
// 	}
// }

// func (this *MedianFinder) FindMedian() float64 {
// 	if this.large.Len() > this.small.Len() {
// 		return float64(this.large[0])
// 	}
// 	if this.small.Len() > this.large.Len() {
// 		return float64(this.small[0])
// 	}
// 	return float64(this.large[0]+this.small[0]) / 2
// }

// /**
//  * Your MedianFinder object will be instantiated and called as such:
//  * obj := Constructor();
//  * obj.AddNum(num);
//  * param_2 := obj.FindMedian();
//  */

// func TestMedianFinder() {
// 	obj := Constructor()
// 	obj.AddNum(40)
// 	obj.AddNum(12)
// 	obj.AddNum(16)
// 	obj.AddNum(14)
// 	obj.AddNum(35)
// 	obj.AddNum(19)
// 	obj.AddNum(34)
// 	obj.AddNum(35)
// 	obj.AddNum(28)
// 	obj.AddNum(35)
// 	obj.AddNum(26)
// 	obj.AddNum(6)
// 	obj.AddNum(8)
// 	obj.AddNum(2)
// 	obj.AddNum(14)
// 	obj.AddNum(25)
// 	obj.AddNum(25)
// 	fmt.Println(obj.FindMedian())
// }
