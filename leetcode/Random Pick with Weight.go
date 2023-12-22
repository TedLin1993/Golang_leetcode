package leetcode

import "math/rand"

type Solution struct {
	preSum []int
}

func Constructor_528(w []int) Solution {
	sum := 0
	preSum := make([]int, len(w)+1)
	preSum[0] = 0
	for i, v := range w {
		sum += v
		preSum[i+1] = sum
	}
	return Solution{preSum}
}

func (this *Solution) PickIndex() int {
	left, right := 0, len(this.preSum)-1
	rnd := rand.Intn(this.preSum[right])
	for left < right {
		mid := left + (right-left+1)/2
		if this.preSum[mid] == rnd {
			return mid
		} else if this.preSum[mid] < rnd {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return left
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */
