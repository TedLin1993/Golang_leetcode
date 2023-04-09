package leetcode

func numEnclaves(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	for i := 0; i < m; i++ {
		dfs_numEnclaves(&grid, i, 0)
		dfs_numEnclaves(&grid, i, n-1)
	}
	for i := 1; i < n-1; i++ {
		dfs_numEnclaves(&grid, 0, i)
		dfs_numEnclaves(&grid, m-1, i)
	}
	sum := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				sum++
			}
		}
	}
	return sum
}

func dfs_numEnclaves(grid *[][]int, row int, column int) {
	if (*grid)[row][column] == 0 {
		return
	}
	(*grid)[row][column] = 0
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for _, dir := range dirs {
		r, c := row+dir[0], column+dir[1]
		if r >= 0 && r < len(*grid) && c >= 0 && c < len((*grid)[0]) {
			dfs_numEnclaves(grid, r, c)
		}
	}
}
