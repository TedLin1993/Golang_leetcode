package leetcode

func numTilings(n int) int {
	dp := make([]int, 1001)
	dp[1] = 1
	dp[2] = 2
	dp[3] = 5
	if n <= 3 {
		return dp[n]
	}
	for i := 4; i <= n; i++ {
		dp[i] = (dp[i-1]*2 + dp[i-3]) % 1000000007
	}
	return dp[n]
}

func TestDominoAndTrominoTiling() {
	// println(numTilings(2))  //2
	// println(numTilings(5))  //24
	println(numTilings(30)) //312342182
}
