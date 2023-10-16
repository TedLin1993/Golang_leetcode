package leetcode

func constructProductMatrix(grid [][]int) [][]int {
	mod := 12345
	n := len(grid)
	m := len(grid[0])

	pre := make([]int, n*m+1)
	pre[0] = 1
	for i := range grid {
		for j := range grid[i] {
			idx := i*m + j
			pre[idx+1] = pre[idx] * grid[i][j] % mod
		}
	}

	suffix := make([]int, n*m+1)
	suffix[n*m] = 1
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			idx := i*m + j
			suffix[idx] = suffix[idx+1] * grid[i][j] % mod
		}
	}

	res := make([][]int, n)
	for i := range grid {
		res[i] = make([]int, m)
		for j := range grid[i] {
			idx := i*m + j
			res[i][j] = pre[idx] * suffix[idx+1] % mod
		}
	}
	return res
}
