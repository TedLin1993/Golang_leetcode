package leetcode

func orangesRotting(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	rottens := make([][]int, 0, m*n)
	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == 2 {
				rottens = append(rottens, []int{r, c})
			}
		}
	}
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	res := 0
	for len(rottens) > 0 {
		length := len(rottens)
		for _, cell := range rottens[:length] {
			for _, dir := range dirs {
				r, c := cell[0]+dir[0], cell[1]+dir[1]
				if r >= 0 && r < m && c >= 0 && c < n && grid[r][c] == 1 {
					grid[r][c] = 2
					rottens = append(rottens, []int{r, c})
				}
			}
		}
		if len(rottens) > length {
			res++
		}
		rottens = rottens[length:]
	}
	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == 1 {
				return -1
			}
		}
	}
	return res
}
