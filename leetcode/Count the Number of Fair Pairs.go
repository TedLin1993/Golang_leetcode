package leetcode

import "sort"

func countFairPairs(nums []int, lower int, upper int) int64 {
	res := 0
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] >= lower && nums[i]+nums[j] <= upper {
				res++
			} else if nums[i]+nums[j] > upper {
				break
			}
		}
	}
	return int64(res)
}
