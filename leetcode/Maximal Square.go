package leetcode

func maximalSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	maxLen := 0
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		if matrix[i][0] == '1' {
			maxLen = 1
			dp[i][0] = 1
		}
	}
	for i := 0; i < n; i++ {
		if matrix[0][i] == '1' {
			maxLen = 1
			dp[0][i] = 1
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
				maxLen = max(maxLen, dp[i][j])
			}
		}
	}
	return maxLen * maxLen
}
