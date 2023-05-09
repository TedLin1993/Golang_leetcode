package leetcode

import "fmt"

func minIncrements(n int, cost []int) int {
	res := 0
	var dfs func(idx int) int
	dfs = func(idx int) int {
		if idx >= n {
			return 0
		}
		left := dfs(2*idx + 1)
		right := dfs(2*idx + 2)
		res += abs(left - right)
		return cost[idx] + max(left, right)
	}
	dfs(0)
	return res
}

func Test_minIncrements() {
	fmt.Println(minIncrements(7, []int{1, 5, 2, 2, 3, 3, 1}))
	fmt.Println(minIncrements(15, []int{764, 1460, 2664, 764, 2725, 4556, 5305, 8829, 5064, 5929, 7660, 6321, 4830, 7055, 3761}))
	fmt.Println(minIncrements(15, []int{764, 1460, 2664, 764, 2725, 4556, 5305, 8829, 5064, 5929, 7660, 6321, 4830, 7055, 3761}))
}
