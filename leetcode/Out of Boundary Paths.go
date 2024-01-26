package leetcode

func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	if maxMove == 0 {
		return 0
	}
	dp := make([][]int, maxMove+1)
	for i := range dp {
		dp[i] = make([]int, m*n)
	}
	for i := 0; i < m*n; i++ {
		r, c := i/n, i%n
		if r == 0 {
			dp[1][i]++
		}
		if r == m-1 {
			dp[1][i]++
		}
		if c == 0 {
			dp[1][i]++
		}
		if c == n-1 {
			dp[1][i]++
		}
	}
	mod := int(1e9 + 7)
	for i := 2; i <= maxMove; i++ {
		for j := 0; j < m*n; j++ {
			r, c := j/n, j%n
			dp[i][j] = dp[1][j]
			if r != 0 {
				dp[i][j] = (dp[i][j] + dp[i-1][j-n]) % mod
			}
			if r != m-1 {
				dp[i][j] = (dp[i][j] + dp[i-1][j+n]) % mod
			}
			if c != 0 {
				dp[i][j] = (dp[i][j] + dp[i-1][j-1]) % mod
			}
			if c != n-1 {
				dp[i][j] = (dp[i][j] + dp[i-1][j+1]) % mod
			}
		}
	}
	return dp[maxMove][startRow*n+startColumn]
}
