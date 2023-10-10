package leetcode

import "strings"

func minOperations_2896(s1 string, s2 string, x int) int {
	if strings.Count(s1, "1")%2 != strings.Count(s2, "1")%2 {
		return -1
	}

	n := len(s1)
	maxVal := n * x
	memo := make([][][2]int, n)
	for i := range memo {
		memo[i] = make([][2]int, n)
		for j := range memo[i] {
			memo[i][j] = [2]int{-1, -1}
		}
	}
	var dfs func(i int, j int, isPre int) int
	dfs = func(i int, j int, isPre int) int {
		if i >= n {
			if j > 0 || isPre == 1 {
				return maxVal
			}
			return 0
		}
		if memo[i][j][isPre] != -1 {
			return memo[i][j][isPre]
		}
		if (s1[i] == s2[i]) == (isPre == 0) {
			memo[i][j][isPre] = dfs(i+1, j, 0)
			return memo[i][j][isPre]
		}
		res := min(dfs(i+1, j+1, 0)+x, dfs(i+1, j, 1)+1)
		if j > 0 {
			res = min(res, dfs(i+1, j-1, 0))
		}
		memo[i][j][isPre] = res
		return res
	}
	return dfs(0, 0, 0)
}
func minOperations_2896_dp(s1 string, s2 string, x int) int {
	if s1 == s2 {
		return 0
	}
	n := len(s1)
	p := make([]int, 0, n)
	for i := range s1 {
		if s1[i] != s2[i] {
			p = append(p, i)
		}
	}
	if len(p)%2 != 0 {
		return -1
	}
	dp0, dp1 := 0, x
	for i := 1; i < len(p); i++ {
		dp0, dp1 = dp1, min(dp1+x, dp0+(p[i]-p[i-1])*2)
	}
	return dp1 / 2
}
