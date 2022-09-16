package leetcode

import (
	"fmt"
)

func maximumScore(nums []int, multipliers []int) int {
	n, m := len(nums), len(multipliers)
	//dp[operation in left count] = score
	dp := make([]int, m+1)

	for op := 1; op <= m; op++ {
		dp[op] = dp[op-1] + multipliers[op-1]*nums[op-1]
		for left := op - 1; left > 0; left-- {
			dp[left] =
				max(dp[left]+multipliers[op-1]*nums[n-op+left],
					dp[left-1]+multipliers[op-1]*nums[left-1])
		}
		dp[0] = dp[0] + multipliers[op-1]*nums[n-op]
	}
	res := dp[0]
	for i := 1; i <= m; i++ {
		res = max(res, dp[i])
	}
	return res
}

func Test_maximumScore() {
	fmt.Println(maximumScore([]int{1, 2, 3}, []int{3, 2, 1}))                       //14
	fmt.Println(maximumScore([]int{-5, -3, -3, -2, 7, 1}, []int{-10, -5, 3, 4, 6})) //102
}
