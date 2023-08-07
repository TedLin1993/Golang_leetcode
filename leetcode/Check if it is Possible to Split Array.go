package leetcode

func canSplitArray(nums []int, m int) bool {
	n := len(nums)
	if n <= 2 {
		return true
	}
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	for i := 0; i < n-1; i++ {
		if nums[i]+nums[i+1] >= m {
			dp[i][i+1] = true
		} else {
			dp[i][i+1] = false
		}
	}
	for gap := 2; gap < n; gap++ {
		for left := 0; left < n-gap; left++ {
			right := left + gap
			if dp[left][right-1] || dp[left+1][right] {
				dp[left][right] = true
			}
		}
	}
	return dp[0][n-1]
}
