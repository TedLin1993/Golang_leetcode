package leetcode

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	res = append(res, []int{})
	cur := []int{}
	var dfs func(idx int)
	dfs = func(idx int) {
		if idx == len(nums) {
			return
		}
		cur = append(cur, nums[idx])
		temp := make([]int, len(cur))
		copy(temp, cur)
		res = append(res, temp)
		dfs(idx + 1)
		cur = cur[:len(cur)-1]
		dfs(idx + 1)
	}
	dfs(0)
	return res
}
