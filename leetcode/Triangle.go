package leetcode

import (
	"fmt"
)

func minimumTotal(triangle [][]int) int {
	dp := make([]int, len(triangle))
	length := len(triangle)
	for i := 0; i < length; i++ {
		dp[i] = triangle[length-1][i]
	}

	for row := len(triangle) - 1; row > 0; row-- {
		dp[0] = dp[0] + triangle[row-1][0]
		lastColumn := row
		for i := 1; i < lastColumn; i++ {
			dp[i-1] = min(dp[i-1], dp[i]+triangle[row-1][i-1])
			dp[i] = dp[i] + triangle[row-1][i]
		}
		dp[lastColumn-1] = min(dp[lastColumn-1], dp[lastColumn]+triangle[row-1][lastColumn-1])
	}
	return dp[0]
}

func TestminimumTotal() {
	fmt.Println(minimumTotal([][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}})) //11
	fmt.Println(minimumTotal([][]int{{-10}}))                                //-10
}
