package leetcode

import (
	"fmt"
)

func maximizeTheProfit(n int, offers [][]int) int {
	group := make([][][]int, n)
	for _, offer := range offers {
		start, end, gold := offer[0], offer[1], offer[2]
		group[end] = append(group[end], []int{start, gold})
	}

	dp := make([]int, n+1)
	for end, g := range group {
		dp[end+1] = dp[end]
		for i := range g {
			start, gold := g[i][0], g[i][1]
			dp[end+1] = max(dp[end+1], dp[start]+gold)
		}
	}
	return dp[n]
}

func Test_maximizeTheProfit() {
	fmt.Println(maximizeTheProfit(5, [][]int{{0, 0, 1}, {0, 2, 2}, {1, 3, 2}}))  //3
	fmt.Println(maximizeTheProfit(5, [][]int{{0, 0, 1}, {0, 2, 10}, {1, 3, 2}})) //10
}
