package leetcode

func mostPoints(questions [][]int) int64 {
	memory := make([]int, len(questions))
	var dfs func(idx int) int
	dfs = func(idx int) int {
		if idx >= len(questions) {
			return 0
		}
		if memory[idx] > 0 {
			return memory[idx]
		}
		solvePoint := questions[idx][0] + dfs(idx+questions[idx][1]+1)
		SkipPoint := dfs(idx + 1)
		memory[idx] = max(solvePoint, SkipPoint)
		return memory[idx]
	}
	return int64(dfs(0))
}

func mostPoints_dp(questions [][]int) int64 {
	dp := make([]int, len(questions))
	res := 0
	for idx, question := range questions {
		if idx > 0 {
			dp[idx] = max(dp[idx], dp[idx-1])
		}
		nextIdx := idx + question[1] + 1
		if nextIdx < len(dp) {
			dp[nextIdx] = max(dp[nextIdx], dp[idx]+question[0])
		} else {
			res = max(res, dp[idx]+question[0])
		}
	}
	return int64(res)
}
