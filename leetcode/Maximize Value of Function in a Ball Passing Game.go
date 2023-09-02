package leetcode

import "math/bits"

func getMaxFunctionValue(receiver []int, k int64) int64 {
	n, m := len(receiver), bits.Len(uint(k))
	type pair struct {
		idx int
		sum int
	}
	dp := make([][]pair, m)
	for i := range dp {
		dp[i] = make([]pair, n)
	}
	for i, r := range receiver {
		dp[0][i] = pair{r, r}
	}
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			idx, sum1 := dp[i-1][j].idx, dp[i-1][j].sum
			nextIdx, sum2 := dp[i-1][idx].idx, dp[i-1][idx].sum
			dp[i][j] = pair{nextIdx, sum1 + sum2}
		}
	}

	res := 0
	for i := 0; i < n; i++ {
		temp := i
		idx := i
		for j := 0; k>>j > 0; j++ {
			if k>>j%2 == 0 {
				continue
			}
			temp += dp[j][idx].sum
			idx = dp[j][idx].idx
		}
		res = max(res, temp)
	}
	return int64(res)
}
