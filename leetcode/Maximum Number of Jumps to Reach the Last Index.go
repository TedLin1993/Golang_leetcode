package leetcode

import "fmt"

func maximumJumps_dfs(nums []int, target int) int {
	n := len(nums)
	targetMap := make([]int, n)
	var dfs func(idx int) int
	dfs = func(idx int) int {
		if idx == n-1 {
			return 0
		}
		res := -1
		for i := idx + 1; i < n; i++ {
			if abs(nums[i]-nums[idx]) <= target {
				if targetMap[i] == 0 {
					dfs(i)
				}
				if targetMap[i] != -1 {
					res = max(res, 1+targetMap[i])
				}
			}
		}
		targetMap[idx] = res
		return res
	}
	dfs(0)
	return targetMap[0]
}

func maximumJumps(nums []int, target int) int {
	n := len(nums)
	dp := make([]int, n)
	for i := 1; i < n; i++ {
		dp[i] = -1
	}
	for i := range dp {
		if dp[i] == -1 {
			continue
		}
		for j := i + 1; j < n; j++ {
			if abs(nums[j]-nums[i]) <= target {
				dp[j] = max(dp[j], dp[i]+1)
			}
		}
	}
	return dp[n-1]
}

func Test_maximumJumps() {
	fmt.Println(maximumJumps([]int{1, 0, 2}, 1))          //1
	fmt.Println(maximumJumps([]int{1, 3, 6, 4, 1, 2}, 2)) //3
	fmt.Println(maximumJumps([]int{1, 3, 6, 4, 1, 2}, 3)) //5
	fmt.Println(maximumJumps([]int{1, 3, 6, 4, 1, 2}, 0)) //-1
}
