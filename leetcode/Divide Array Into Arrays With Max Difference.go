package leetcode

import "slices"

func divideArray(nums []int, k int) [][]int {
	slices.Sort(nums)
	n := len(nums)
	res := make([][]int, 0, n/3)
	for i := 0; i+2 < n; i += 3 {
		if nums[i+2]-nums[i] > k {
			return nil
		}
		res = append(res, nums[i:i+3])
	}
	return res
}
