package leetcode

func PredictTheWinner(nums []int) bool {
	var maxDiff func(nums []int) int
	maxDiff = func(nums []int) int {
		n := len(nums)
		if n == 1 {
			return nums[0]
		}
		return max(nums[0]-maxDiff(nums[1:]), nums[n-1]-maxDiff(nums[:n-1]))
	}
	return maxDiff(nums) >= 0
}

func PredictTheWinner_dp(nums []int) bool {
	n := len(nums)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = nums[i]
	}
	for diff := 1; diff < n; diff++ {
		for left := 0; left < n-diff; left++ {
			right := left + diff
			dp[left][right] = max(nums[left]-dp[left+1][right], nums[right]-dp[left][right-1])
		}
	}
	return dp[0][n-1] >= 0
}
