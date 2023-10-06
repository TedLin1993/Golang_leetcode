package leetcode

func integerBreak(n int) int {
	if n < 4 {
		return n - 1
	}
	dp := make([]int, n+1)
	dp[2] = 2
	dp[3] = 3
	for i := 4; i <= n; i++ {
		for j := 1; j < i; j++ {
			dp[i] = max(dp[i], dp[i-j]*dp[j])
		}
	}
	return dp[n]
}
