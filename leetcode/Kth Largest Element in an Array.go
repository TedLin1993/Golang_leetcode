package leetcode

import (
	"fmt"
	"math/rand"
)

func findKthLargest(nums []int, k int) int {

}

func kthLargest(nums []int, targeIdx int, left int, right int) int {
	pivotIdx := rand.Intn(right-left+1) + left
	pivot := nums[pivotIdx]

	//pivot move to right
	nums[pivotIdx], nums[right] = nums[right], nums[pivotIdx]

	pivotIdx = left
	for i := left + 1; i < right; i++ {
		if nums[i] < pivot {
			nums[pivotIdx], nums[i] = nums[i], nums[pivotIdx]
			pivotIdx++
		}
	}
	//TODO
}

func TestfindKthLargest() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))          //5
	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4)) //4
}
