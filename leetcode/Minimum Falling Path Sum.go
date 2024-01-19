package leetcode

import "slices"

func minFallingPathSum(m [][]int) int {
	n := len(m)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		dp[n-1][i] = m[n-1][i]
	}
	for i := n - 2; i >= 0; i-- {
		dp[i][0] = m[i][0] + min(dp[i+1][0], dp[i+1][1])
		for j := 1; j < n-1; j++ {
			dp[i][j] = m[i][j] + min(dp[i+1][j-1], dp[i+1][j], dp[i+1][j+1])
		}
		dp[i][n-1] = m[i][n-1] + min(dp[i+1][n-2], dp[i+1][n-1])
	}
	return slices.Min(dp[0])
}
