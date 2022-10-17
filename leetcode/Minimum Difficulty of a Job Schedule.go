package leetcode

import (
	"fmt"
	"math"
)

func minDifficulty(jobDifficulty []int, d int) int {
	if len(jobDifficulty) < d {
		return -1
	}

	n := len(jobDifficulty)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, d)
		for j := 0; j < d; j++ {
			dp[i][j] = math.MaxInt16
		}
	}

	dp[0][0] = jobDifficulty[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], jobDifficulty[i])
	}

	for i := 1; i < d; i++ {
		for j := 0; j < n; j++ {
			for k := i - 1; k < j; k++ {
				dp[j][i] = min(dp[j][i], dp[k][i-1]+maxArr(jobDifficulty[k+1:j+1]))
			}
		}
	}
	return dp[n-1][d-1]
}

func Test_minDifficulty() {
	fmt.Println(minDifficulty([]int{6, 5, 4, 3, 2, 1}, 2)) //7
	fmt.Println(minDifficulty([]int{9, 9, 9}, 4))          //-1
	fmt.Println(minDifficulty([]int{1, 1, 1}, 3))          //3
}
