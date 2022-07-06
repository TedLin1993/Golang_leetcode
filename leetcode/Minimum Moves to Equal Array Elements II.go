package leetcode

import (
	"fmt"
	"math/rand"
	"time"
)

func minMoves2(nums []int) int {
	rand.Seed(time.Now().UnixNano())
	mid := getMid(nums, 0, len(nums)-1)
	res := 0
	for _, n := range nums {
		diff := mid - n
		if diff < 0 {
			res -= diff
		} else {
			res += diff
		}
	}
	return res
}

func getMid(nums []int, left, right int) int {
	pivotIdx := left + rand.Intn(right-left+1)
	pivot := nums[pivotIdx]

	nums[pivotIdx], nums[right] = nums[right], nums[pivotIdx]
	pivotIdx = left
	for i := left; i < right; i++ {
		if nums[i] < pivot {
			nums[pivotIdx], nums[i] = nums[i], nums[pivotIdx]
			pivotIdx++
		}
	}
	nums[pivotIdx], nums[right] = nums[right], nums[pivotIdx]

	if pivotIdx == len(nums)/2 {
		return pivot
	}
	if pivotIdx < len(nums)/2 {
		return getMid(nums, pivotIdx+1, right)
	}
	return getMid(nums, left, pivotIdx-1)
}

func TestminMoves2() {
	fmt.Println(minMoves2([]int{1, 2, 3}))     //2
	fmt.Println(minMoves2([]int{1, 10, 2, 9})) //16
	fmt.Println(minMoves2([]int{1, 1, 1}))     //0
}
