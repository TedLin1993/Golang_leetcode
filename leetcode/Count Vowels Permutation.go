package leetcode

func countVowelPermutation(n int) int {
	dp := make([]int, 5)

	//a, e , i, o ,u represent dp[0],dp[1],dp[2],dp[3],dp[4]
	//when n=1, all element is 1
	for i := 0; i < 5; i++ {
		dp[i] = 1
	}

	mod := 1000000007
	for i := 1; i < n; i++ {
		nextDp := make([]int, 5)
		nextDp[0] = (dp[1] + dp[2] + dp[4]) % mod
		nextDp[1] = (dp[0] + dp[2]) % mod
		nextDp[2] = (dp[1] + dp[3]) % mod
		nextDp[3] = dp[2] % mod
		nextDp[4] = (dp[2] + dp[3]) % mod
		dp = nextDp
	}
	res := 0
	for i := 0; i < len(dp); i++ {
		res += dp[i]
	}
	return res % mod
}
