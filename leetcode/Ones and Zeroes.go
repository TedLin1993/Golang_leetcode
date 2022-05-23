package leetcode

import (
	"fmt"
)

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	for _, str := range strs {
		zeroCount := 0
		oneCount := 0
		for i := 0; i < len(str); i++ {
			if str[i] == '0' {
				zeroCount++
			} else {
				oneCount++
			}
		}

		for i := m; i >= zeroCount; i-- {
			for j := n; j >= oneCount; j-- {
				dp[i][j] = max(dp[i][j], dp[i-zeroCount][j-oneCount]+1)
			}
		}
	}
	return dp[m][n]
}

func TestfindMaxForm() {
	fmt.Println(findMaxForm([]string{"10", "0001", "111001", "1", "0"}, 5, 3)) //4
	fmt.Println(findMaxForm([]string{"10", "1", "0"}, 1, 1))                   //2
}
