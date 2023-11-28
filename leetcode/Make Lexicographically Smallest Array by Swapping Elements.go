package leetcode

import "sort"

func lexicographicallySmallestArray(nums []int, limit int) []int {
	n := len(nums)
	type pair struct {
		idx int
		num int
	}
	pairs := make([]pair, n)
	for i, v := range nums {
		pairs[i] = pair{idx: i, num: v}
	}
	sort.Slice(pairs, func(a, b int) bool {
		return pairs[a].num < pairs[b].num
	})
	res := make([]int, n)
	left := 0
	for left < n {
		right := left
		for i := left + 1; i < n; i++ {
			if pairs[i].num-pairs[i-1].num <= limit {
				right = i
			} else {
				break
			}
		}
		ids := make([]int, 0, right-left+1)
		for i := left; i <= right; i++ {
			ids = append(ids, pairs[i].idx)
		}
		sort.Ints(ids)
		for i, idx := range ids {
			res[idx] = pairs[left+i].num
		}
		left = right + 1
	}
	return res
}
