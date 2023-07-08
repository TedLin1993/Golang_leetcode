package leetcode

import "sort"

func putMarbles(weights []int, k int) int64 {
	n := len(weights)
	pairs := make([]int, n-1)
	for i := 1; i < n; i++ {
		pairs[i-1] = weights[i-1] + weights[i]
	}
	sort.Ints(pairs)
	max, min := 0, 0
	for i := 0; i < k-1; i++ {
		max += pairs[n-2-i]
		min += pairs[i]
	}
	return int64(max - min)
}
