package leetcode

import "fmt"

func maxAreaOfIsland(grid [][]int) int {
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				res = max(res, dfsAreaOfIsland(grid, i, j))
			}
		}
	}
	return res
}

func dfsAreaOfIsland(grid [][]int, row, column int) int {
	if row < 0 || row >= len(grid) || column < 0 || column >= len(grid[0]) || grid[row][column] == 0 {
		return 0
	}
	res := 1
	grid[row][column] = 0
	res += dfsAreaOfIsland(grid, row, column-1)
	res += dfsAreaOfIsland(grid, row, column+1)
	res += dfsAreaOfIsland(grid, row-1, column)
	res += dfsAreaOfIsland(grid, row+1, column)
	return res
}

func TestmaxAreaOfIsland() {
	fmt.Println(maxAreaOfIsland([][]int{{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}}))
}
