package leetcode

func longestCommonSubsequence(text1 string, text2 string) int {
	if len(text1) < len(text2) {
		text1, text2 = text2, text1
	}
	n := len(text2)
	dp := make([]int, n+1)
	for _, c1 := range text1 {
		temp := make([]int, n+1)
		for i, c2 := range text2 {
			if c1 == c2 {
				temp[i+1] = dp[i] + 1
			} else {
				temp[i+1] = max(dp[i+1], temp[i])
			}
		}
		dp = temp
	}
	return dp[n]
}
