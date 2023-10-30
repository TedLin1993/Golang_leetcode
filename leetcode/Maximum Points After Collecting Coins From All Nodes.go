package leetcode

func maximumPoints(edges [][]int, coins []int, k int) int {
	n := len(coins)
	g := make([][]int, n)
	for _, e := range edges {
		g[e[0]] = append(g[e[0]], e[1])
		g[e[1]] = append(g[e[1]], e[0])
	}
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, 14)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dfs func(i int, j int, parent int) int
	dfs = func(i int, j int, parent int) int {
		if memo[i][j] != -1 {
			return memo[i][j]
		}
		res1 := coins[i]>>j - k
		res2 := coins[i] >> (j + 1)
		for _, next := range g[i] {
			if next == parent {
				continue
			}
			res1 += dfs(next, j, i)
			if j < 13 {
				res2 += dfs(next, j+1, i)
			}
		}
		memo[i][j] = max(res1, res2)
		return memo[i][j]
	}
	return dfs(0, 0, -1)
}
