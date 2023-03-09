package leetcode

import (
	"fmt"
)

func waysToReachTarget(target int, types [][]int) int {
	mod := int(1e9 + 7)
	dp := make([]int, target+1)
	dp[0] = 1
	for _, t := range types {
		count, mark := t[0], t[1]
		for currentTarget := target; currentTarget >= mark; currentTarget-- {
			for i := 1; i <= count && i*mark <= currentTarget; i++ {
				dp[currentTarget] += dp[currentTarget-i*mark]
			}
			dp[currentTarget] %= mod
		}
	}
	return dp[target]
}

func Test_waysToReachTarget() {
	fmt.Println(waysToReachTarget(6, [][]int{{6, 1}, {3, 2}, {2, 3}}))    //7
	fmt.Println(waysToReachTarget(5, [][]int{{50, 1}, {50, 2}, {50, 5}})) //4
	fmt.Println(waysToReachTarget(18, [][]int{{6, 1}, {3, 2}, {2, 3}}))   //1
}
