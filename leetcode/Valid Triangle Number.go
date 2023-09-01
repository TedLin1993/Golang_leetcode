package leetcode

import "sort"

func triangleNumber(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	res := 0
	for k := n - 1; k > 1; k-- {
		i, j := 0, k-1
		for i < j {
			if nums[i]+nums[j] > nums[k] {
				res += j - i
				j--
			} else {
				i++
			}
		}
	}
	return res
}
