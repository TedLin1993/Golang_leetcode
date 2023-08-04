package leetcode

func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	dp := make([]bool, n)
	for i := 0; i < n; i++ {
		if i > 0 && !dp[i-1] {
			continue
		}
		for _, w := range wordDict {
			m := len(w)
			if i+m-1 < n && s[i:i+m] == w {
				dp[i+m-1] = true
			}
		}
	}
	return dp[n-1]
}
