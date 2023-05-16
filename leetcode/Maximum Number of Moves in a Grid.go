package leetcode

func maxMoves(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	memory := make([][]int, m)
	for i := 0; i < m; i++ {
		memory[i] = make([]int, n)
		for j := 0; j < n; j++ {
			memory[i][j] = -1
		}
	}
	var dfs func(r, c int) int
	dfs = func(r, c int) int {
		if c == n-1 {
			return 1
		}
		if memory[r][c] != -1 {
			return memory[r][c]
		}
		res := 1
		if r-1 >= 0 && grid[r][c] < grid[r-1][c+1] {
			res = max(res, 1+dfs(r-1, c+1))
		}
		if grid[r][c] < grid[r][c+1] {
			res = max(res, 1+dfs(r, c+1))
		}
		if r+1 < m && grid[r][c] < grid[r+1][c+1] {
			res = max(res, 1+dfs(r+1, c+1))
		}
		memory[r][c] = res
		return res
	}
	res := 0
	for i := 0; i < m; i++ {
		res = max(res, dfs(i, 0))
	}
	return res - 1
}

func maxMoves_bfs(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	queue := make([][]int, 0, m)
	for i := 0; i < m; i++ {
		queue = append(queue, []int{i, 0})
	}
	res := 0
	directions := [][]int{{-1, 1}, {0, 1}, {1, 1}}
	for len(queue) > 0 {
		r, c := queue[0][0], queue[0][1]
		if c == n-1 {
			return c
		}
		if c > res {
			res = c
		}
		for _, dir := range directions {
			rNew, cNew := r+dir[0], c+dir[1]
			if rNew >= 0 && rNew < m && grid[r][c] < grid[rNew][cNew] && !visited[rNew][cNew] {
				visited[rNew][cNew] = true
				queue = append(queue, []int{rNew, cNew})
			}
		}
		queue = queue[1:]
	}
	return res
}
