package leetcode

func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	set := make([]int, m*n)
	for i := range set {
		r := i / n
		c := i % n
		if grid[r][c] == '1' {
			set[i] = i
		} else {
			set[i] = -1
		}
	}
	var find func(idx int) int
	find = func(idx int) int {
		if set[idx] == idx {
			return idx
		}
		if set[idx] == -1 {
			return -1
		}
		set[idx] = find(set[idx])
		return set[idx]
	}
	union := func(a, b int) {
		a, b = find(a), find(b)
		if a > b {
			a, b = b, a
		}
		set[b] = a
	}
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '0' {
				continue
			}
			idx1 := i*n + j
			for _, dir := range dirs {
				r, c := i+dir[0], j+dir[1]
				if r >= 0 && r < m && c >= 0 && c < n && grid[r][c] == '1' {
					idx2 := r*n + c
					union(idx1, idx2)
				}
			}
		}
	}

	res := 0
	for i := range set {
		if find(i) == i {
			res++
		}
	}
	return res
}
