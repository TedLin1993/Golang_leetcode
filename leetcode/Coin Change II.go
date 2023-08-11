package leetcode

func change(amount int, coins []int) int {
	n := len(coins)
	memo := make([][]int, amount+1)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dfs func(amount int, idx int) int
	dfs = func(amount int, idx int) int {
		if amount == 0 {
			return 1
		}
		if idx == n {
			return 0
		}
		if memo[amount][idx] != -1 {
			return memo[amount][idx]
		}
		res := 0
		for i := 0; amount-coins[idx]*i >= 0; i++ {
			res += dfs(amount-coins[idx]*i, idx+1)
		}
		memo[amount][idx] = res
		return res
	}
	return dfs(amount, 0)
}

func change_dp(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, c := range coins {
		for i := c; i <= amount; i++ {
			dp[i] += dp[i-c]
		}
	}
	return dp[amount]
}
