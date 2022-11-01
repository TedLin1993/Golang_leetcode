package leetcode

import "fmt"

func findBall(grid [][]int) []int {
	colLen := len(grid[0])
	dp := make([]int, colLen)
	for i := 0; i < colLen; i++ {
		dp[i] = i
	}

	for rowIdx := 0; rowIdx < len(grid); rowIdx++ {
		for i := 0; i < colLen; i++ {
			if dp[i] == -1 {
				continue
			}
			pos := dp[i]
			if grid[rowIdx][pos] == -1 {
				if pos == 0 || grid[rowIdx][pos-1] == 1 {
					dp[i] = -1
				} else {
					dp[i]--
				}
			} else {
				if pos == colLen-1 || grid[rowIdx][pos+1] == -1 {
					dp[i] = -1
				} else {
					dp[i]++
				}
			}
		}
	}
	return dp
}

func Test_findBall() {
	// fmt.Println(findBall([][]int{{1, 1, 1, -1, -1}, {1, 1, 1, -1, -1}, {-1, -1, -1, 1, 1}, {1, 1, 1, 1, -1}, {-1, -1, -1, -1, -1}}))                                                                                                                                                                                                                                                                                                                                                                //[1 -1 -1 -1 -1]
	// fmt.Println(findBall([][]int{{-1}}))                                                                                                                                                                                                                                                                                                                                                                                                                                                            //[-1]
	// fmt.Println(findBall([][]int{{1, 1, 1, 1, 1, 1}, {-1, -1, -1, -1, -1, -1}, {1, 1, 1, 1, 1, 1}, {-1, -1, -1, -1, -1, -1}}))                                                                                                                                                                                                                                                                                                                                                                      //[0,1,2,3,4,-1]
	fmt.Println(findBall([][]int{{1, -1, -1, 1, -1, 1, 1, 1, 1, 1, -1, 1, 1, 1, 1, 1, 1, -1, -1, -1, -1, -1, -1, 1, -1, 1, -1, 1, -1, -1, -1, -1, 1, -1, 1, 1, -1, -1, -1, -1, -1, 1}, {-1, 1, 1, 1, -1, -1, -1, -1, 1, 1, 1, -1, -1, -1, 1, -1, -1, 1, 1, 1, 1, 1, 1, -1, 1, -1, -1, -1, -1, -1, 1, -1, 1, -1, -1, -1, -1, 1, 1, -1, 1, 1}, {1, -1, -1, -1, -1, 1, -1, 1, 1, 1, 1, 1, 1, 1, -1, 1, -1, -1, -1, 1, -1, -1, 1, -1, 1, -1, 1, -1, -1, 1, -1, 1, -1, 1, 1, -1, -1, 1, 1, -1, 1, -1}})) //[-1]
}
