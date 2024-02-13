package leetcode

import "slices"

func largestDivisibleSubset(nums []int) []int {
	slices.Sort(nums)
	n := len(nums)
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}
	prev := make([]int, n)
	idx := 0
	for i := range nums {
		for j := i - 1; j >= 0; j-- {
			if nums[i]%nums[j] == 0 && dp[i] <= dp[j] {
				dp[i] = dp[j] + 1
				prev[i] = j
			}
		}
		if dp[i] > dp[idx] {
			idx = i
		}
	}
	res := make([]int, dp[idx])
	for i := 0; i < len(res); i++ {
		res[i] = nums[idx]
		idx = prev[idx]
	}
	return res
}