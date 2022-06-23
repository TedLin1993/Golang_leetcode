package leetcode

import (
	"fmt"
	"math/rand"
)

func findKthLargest(nums []int, k int) int {
	return kthLargest(nums, len(nums)-k, 0, len(nums)-1)
}

func kthLargest(nums []int, targetIdx int, left int, right int) int {
	pivotIdx := rand.Intn(right-left+1) + left
	pivot := nums[pivotIdx]

	//pivot move to right
	nums[pivotIdx], nums[right] = nums[right], nums[pivotIdx]

	//lower than pivot move to pivotIdx's left, others move to pivotIdx's right
	pivotIdx = left
	for i := left; i < right; i++ {
		if nums[i] <= pivot {
			nums[pivotIdx], nums[i] = nums[i], nums[pivotIdx]
			pivotIdx++
		}
	}

	//pivot move to pivotIdx
	nums[pivotIdx], nums[right] = nums[right], nums[pivotIdx]

	if pivotIdx == targetIdx {
		return nums[pivotIdx]
	} else if pivotIdx < targetIdx {
		left = pivotIdx + 1
		return kthLargest(nums, targetIdx, left, right)
	} else {
		right = pivotIdx - 1
		return kthLargest(nums, targetIdx, left, right)
	}
}

func TestfindKthLargest() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))          //5
	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4)) //4
}
