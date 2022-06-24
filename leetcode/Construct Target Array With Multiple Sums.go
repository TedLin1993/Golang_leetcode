package leetcode

import (
	"container/heap"
	"fmt"
)

func isPossible(target []int) bool {
	if len(target) == 1 {
		if target[0] == 1 {
			return true
		} else {
			return false
		}
	}

	maxHeap := MaxHeap{}
	sum := 0
	for _, t := range target {
		heap.Push(&maxHeap, t)
		sum += t
	}

	for {
		max := heap.Pop(&maxHeap).(int)
		sum -= max
		
		if max == 1 || sum == 1 {
			return true
		}
		if max < sum || max%sum == 0 {
			return false
		}

		max = max % sum
		sum += max
		heap.Push(&maxHeap, max)
	}
}

func TestisPossible() {
	fmt.Println(isPossible([]int{9, 3, 5}))       //true
	fmt.Println(isPossible([]int{1, 1, 1, 2}))    //false
	fmt.Println(isPossible([]int{8, 5}))          //true
	fmt.Println(isPossible([]int{1, 1000000000})) //true
	fmt.Println(isPossible([]int{9, 9, 9}))       //false
	fmt.Println(isPossible([]int{2}))             //false
}
