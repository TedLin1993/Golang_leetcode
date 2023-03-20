package leetcode

func checkValidGrid(grid [][]int) bool {
	n := len(grid)
	count := n * n
	gridMap := make([][]int, count)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			val := grid[i][j]
			gridMap[val] = []int{i, j}
		}
	}

	current := []int{0, 0}
	for i := 1; i < count; i++ {
		cell := gridMap[i]
		x_diff, y_diff := cell[0]-current[0], cell[1]-current[1]
		if x_diff < 0 {
			x_diff = -x_diff
		}
		if y_diff < 0 {
			y_diff = -y_diff
		}

		if !((x_diff == 2 && y_diff == 1) || (x_diff == 1 && y_diff == 2)) {
			return false
		}
		current = cell
	}
	return true
}
