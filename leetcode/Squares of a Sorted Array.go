package leetcode

import "slices"

func sortedSquares(nums []int) []int {
	n := len(nums)
	left, right := 0, n-1
	res := make([]int, 0, n)
	for left <= right {
		if abs(nums[left]) >= abs(nums[right]) {
			res = append(res, nums[left]*nums[left])
			left++
		} else {
			res = append(res, nums[right]*nums[right])
			right--
		}
	}
	slices.Reverse(res)
	return res
}
