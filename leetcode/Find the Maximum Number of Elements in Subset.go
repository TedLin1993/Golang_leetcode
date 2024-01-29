package leetcode

import "slices"

func maximumLength(nums []int) int {
	res := 1
	n := len(nums)
	numMap := make(map[int]int, n)
	for _, v := range nums {
		numMap[v]++
	}
	for numMap[1] >= 3 {
		res += 2
		numMap[1] -= 2
	}
	numMap[1] = 0
	slices.Sort(nums)
	for _, v := range nums {
		if numMap[v] == 0 {
			continue
		}
		temp := 1
		cur := v
		for numMap[cur] >= 2 && numMap[cur*cur] >= 1 {
			numMap[cur] = 0
			temp += 2
			cur *= cur
		}
		res = max(res, temp)
	}
	return res
}
