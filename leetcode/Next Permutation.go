package leetcode

import "sort"

func nextPermutation(nums []int) {
	n := len(nums)
	idx := -1
	for i := n - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			idx = i - 1
			break
		}
	}
	if idx == -1 {
		sort.Ints(nums)
		return
	}

	target := idx + 1
	for i := idx + 2; i < n; i++ {
		if nums[i] > nums[idx] && nums[i] < nums[target] {
			target = i
		}
	}
	nums[idx], nums[target] = nums[target], nums[idx]
	sort.Ints(nums[idx+1:])
}
