package leetcode

func numDecodings(s string) int {
	n := len(s)
	dp := make([]int, n+1)
	dp[n] = 1
	if s[n-1] != '0' {
		dp[n-1] = 1
	}
	for i := n - 2; i >= 0; i-- {
		if s[i] == '0' {
			dp[i] = 0
		} else if s[i] == '1' || (s[i] == '2' && s[i+1] <= '6') {
			dp[i] = dp[i+1] + dp[i+2]
		} else {
			dp[i] = dp[i+1]
		}

	}
	return dp[0]
}

func numDecodings_2(s string) int {
	n := len(s)
	dp0, dp1 := 0, 1
	if s[n-1] != '0' {
		dp0 = 1
	}

	for i := n - 2; i >= 0; i-- {
		if s[i] == '0' {
			dp0, dp1 = 0, dp0
		} else if s[i] == '1' || (s[i] == '2' && s[i+1] <= '6') {
			dp0, dp1 = dp0+dp1, dp0
		} else {
			dp0, dp1 = dp0, dp0
		}
	}
	return dp0
}
