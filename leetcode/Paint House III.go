package leetcode

import (
	"fmt"
	"math"
)

func minCost(houses []int, cost [][]int, m int, n int, target int) int {
	dp := make([][][]int, m+1) //[house][color][target]
	for i := 0; i < m+1; i++ {
		dp[i] = make([][]int, n+1)
		for j := 0; j < n+1; j++ {
			dp[i][j] = make([]int, target+1)
			for k := 0; k < target+1; k++ {
				dp[i][j][k] = math.MaxInt32
			}
		}
	}
	for i := 0; i < n+1; i++ {
		dp[0][i][0] = 0
	}

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			for k := 1; k < target+1; k++ {
				if houses[i-1] != 0 {
					if houses[i-1] != j {
						continue
					}
					for c := 1; c < n+1; c++ {
						if c == j {
							dp[i][j][k] = min(dp[i-1][c][k], dp[i][j][k])
						} else {
							dp[i][j][k] = min(dp[i-1][c][k-1], dp[i][j][k])
						}
					}
				} else {
					for c := 1; c < n+1; c++ {
						if c == j {
							dp[i][j][k] = min(dp[i-1][c][k]+cost[i-1][j-1], dp[i][j][k])
						} else {
							dp[i][j][k] = min(dp[i-1][c][k-1]+cost[i-1][j-1], dp[i][j][k])
						}
					}
				}
			}
		}
	}

	res := math.MaxInt32
	for i := 1; i < n+1; i++ {
		if dp[m][i][target] < res {
			res = dp[m][i][target]
		}
	}
	if res < math.MaxInt32 {
		return res
	} else {
		return -1
	}
}

func TestminCost() {
	fmt.Println(minCost([]int{0, 0, 0, 0, 0}, [][]int{{1, 10}, {10, 1}, {10, 1}, {1, 10}, {5, 1}}, 5, 2, 3)) //9
	fmt.Println(minCost([]int{0, 2, 1, 2, 0}, [][]int{{1, 10}, {10, 1}, {10, 1}, {1, 10}, {5, 1}}, 5, 2, 3)) //11
	fmt.Println(minCost([]int{3, 1, 2, 3}, [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}}, 4, 3, 3))    //-1
}
