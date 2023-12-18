package leetcode

import "slices"

func maxFrequencyScore(nums []int, k int64) int {
	slices.Sort(nums)
	n := len(nums)
	preSum := make([]int, n+1)
	for i, v := range nums {
		preSum[i+1] = preSum[i] + v
	}
	getDistance := func(left, right int) int64 {
		mid := (left + right) / 2
		leftSum := nums[mid]*(mid-left) - (preSum[mid] - preSum[left])
		rightSum := preSum[right+1] - preSum[mid+1] - nums[mid]*(right-mid)
		return int64(leftSum + rightSum)
	}
	left := 0
	res := 0
	for right := range nums {
		for getDistance(left, right) > k {
			left++
		}
		res = max(res, right-left+1)
	}
	return res
}
