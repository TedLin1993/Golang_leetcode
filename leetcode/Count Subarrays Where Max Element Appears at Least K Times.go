package leetcode

import (
	"fmt"
	"slices"
)

func countSubarrays(nums []int, k int) int64 {
	maxVal := slices.Max(nums)
	count := 0
	left := 0
	res := 0
	for _, v := range nums {
		if v == maxVal {
			count++
		}
		for count == k {
			if nums[left] == maxVal {
				count--
			}
			left++
		}
		res += left
	}
	return int64(res)
}

func Test_countSubarrays() {
	fmt.Println(countSubarrays([]int{61, 23, 38, 23, 56, 40, 82, 56, 82, 82, 82, 70, 8, 69, 8, 7, 19, 14, 58, 42, 82, 10, 82, 78, 15, 82}, 2))
}
