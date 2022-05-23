package leetcode

import "math"

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		dp[i] = math.MaxInt32
	}
	for i := 1; i <= amount; i++ {
		for _, c := range coins {
			if i-c >= 0 {
				dp[i] = min(dp[i], dp[i-c]+1)
			}
		}
	}
	if dp[amount] >= math.MaxInt32 {
		return -1
	}
	return dp[amount]
}
