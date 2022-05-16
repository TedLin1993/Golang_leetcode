package leetcode

import (
	"fmt"
)

//dfs, robot#1的每個行動都需要分別計算robot#2的3個行動

func cherryPickup(grid [][]int) int {
	columnLen := len(grid[0])
	rowLen := len(grid)
	dp := make([][][]int, rowLen)
	for i := 0; i < rowLen; i++ {
		dp[i] = make([][]int, columnLen)
		for j := 0; j < columnLen; j++ {
			dp[i][j] = make([]int, columnLen)
		}
	}

	dp[0][0][columnLen-1] = grid[0][0] + grid[0][columnLen-1]
	res := 0
	for row := 1; row < rowLen; row++ {
		for r1col := 0; r1col < columnLen; r1col++ {
			for r2col := 0; r2col < columnLen; r2col++ {
				temp := maxScore(grid, dp, row, r1col, r2col)
				dp[row][r1col][r2col] = temp
				if res < temp {
					res = temp
				}
			}
		}
	}
	return res
}

func maxScore(grid [][]int, dp [][][]int, row, r1col, r2col int) int {
	columnLen := len(grid[0])
	res := 0
	for i := -1; i < 2; i++ {
		if r1col+i < 0 || r1col+i >= columnLen {
			continue
		}
		for j := -1; j < 2; j++ {
			if r2col+j < 0 || r2col+j >= columnLen {
				continue
			}

			temp := dp[row-1][r1col+i][r2col+j] + grid[row][r1col]
			if r1col != r2col {
				temp += grid[row][r2col]
			}

			if res < temp {
				res = temp
			}
		}
	}
	return res
}

func TestCherry_Pickup_II() {
	//24
	// fmt.Println(cherryPickup([][]int{{3, 1, 1}, {2, 5, 1}, {1, 5, 5}, {2, 1, 1}}))
	//28
	// fmt.Println(cherryPickup([][]int{{1, 0, 0, 0, 0, 0, 1}, {2, 0, 0, 0, 0, 3, 0}, {2, 0, 9, 0, 0, 0, 0}, {0, 3, 0, 5, 4, 0, 0}, {1, 0, 2, 3, 0, 0, 6}}))
	//96
	fmt.Println(cherryPickup([][]int{{0, 8, 7, 10, 9, 10, 0, 9, 6}, {8, 7, 10, 8, 7, 4, 9, 6, 10}, {8, 1, 1, 5, 1, 5, 5, 1, 2}, {9, 4, 10, 8, 8, 1, 9, 5, 0}, {4, 3, 6, 10, 9, 2, 4, 8, 10}, {7, 3, 2, 8, 3, 3, 5, 9, 8}, {1, 2, 6, 5, 6, 2, 0, 10, 0}}))

}
