package leetcode

import "sort"

func numberOfPoints(nums [][]int) int {
	sort.Slice(nums, func(a, b int) bool {
		return nums[a][0] < nums[b][0]
	})
	res := 0
	cur := nums[0]
	for _, num := range nums[1:] {
		if num[0] <= cur[1] {
			cur[1] = max(cur[1], num[1])
		} else {
			res += cur[1] - cur[0] + 1
			cur = num
		}
	}
	res += cur[1] - cur[0] + 1
	return res
}
