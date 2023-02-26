package leetcode

import (
	"fmt"
	"sort"
)

func maxNumOfMarkedIndices(nums []int) int {
	res := 0
	sort.Ints(nums)
	mid := len(nums) / 2
	left, right := 0, mid
	for left < mid && right < len(nums) {
		if 2*nums[left] <= nums[right] {
			res += 2
			left++
		}
		right++
	}
	return res
}

func Test_maxNumOfMarkedIndices() {
	fmt.Println(maxNumOfMarkedIndices([]int{9, 2, 5, 4}))                                                                                                                                                                                                                             //4
	fmt.Println(maxNumOfMarkedIndices([]int{57, 40, 57, 51, 90, 51, 68, 100, 24, 39, 11, 85, 2, 22, 67, 29, 74, 82, 10, 96, 14, 35, 25, 76, 26, 54, 29, 44, 63, 49, 73, 50, 95, 89, 43, 62, 24, 88, 88, 36, 6, 16, 14, 2, 42, 42, 60, 25, 4, 58, 23, 22, 27, 26, 3, 79, 64, 20, 92})) //58
}
