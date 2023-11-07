package leetcode

func maximumScoreAfterOperations(edges [][]int, values []int) int64 {
	n := len(values)
	g := make([][]int, n)
	for _, e := range edges {
		g[e[0]] = append(g[e[0]], e[1])
		g[e[1]] = append(g[e[1]], e[0])
	}
	var dfs func(cur int, parent int) int
	dfs = func(cur int, parent int) int {
		if len(g[cur]) == 1 && parent != -1 {
			return values[cur]
		}
		res1 := values[cur]
		res2 := 0
		for _, next := range g[cur] {
			if next != parent {
				res2 += dfs(next, cur)
			}
		}
		return min(res1, res2)
	}
	sum := 0
	for _, v := range values {
		sum += v
	}
	return int64(sum - dfs(0, -1))
}
