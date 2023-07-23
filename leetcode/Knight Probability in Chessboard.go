package leetcode

import (
	"fmt"
)

func knightProbability(n int, k int, row int, column int) float64 {
	dirs := [][]int{{1, 2}, {2, 1}, {2, -1}, {1, -2}, {-1, -2}, {-2, -1}, {-2, 1}, {-1, 2}}
	dp := make([][][]float64, k+1)
	for i := range dp {
		dp[i] = make([][]float64, n)
		for j := range dp[i] {
			dp[i][j] = make([]float64, n)
		}
	}
	dp[0][row][column] = 1
	for i := 0; i < k; i++ {
		for r := range dp[i] {
			for c := range dp[i][r] {
				for _, dir := range dirs {
					rNext, cNext := r+dir[0], c+dir[1]
					if rNext >= 0 && rNext < n && cNext >= 0 && cNext < n {
						dp[i+1][rNext][cNext] += dp[i][r][c] / 8
					}
				}
			}
		}
	}
	res := 0.0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			res += dp[k][i][j]
		}
	}
	return res
}

func Test_knightProbability() {
	fmt.Println(knightProbability(3, 2, 0, 0))  //0.0625
	fmt.Println(knightProbability(1, 0, 0, 0))  //1
	fmt.Println(knightProbability(8, 30, 6, 4)) //0.00019
}
