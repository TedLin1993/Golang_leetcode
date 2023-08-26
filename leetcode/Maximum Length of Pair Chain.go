package leetcode

import "sort"

func findLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(a, b int) bool {
		return pairs[a][1] < pairs[b][1]
	})
	res := 1
	right := pairs[0][1]
	for _, p := range pairs[1:] {
		if p[0] > right {
			res++
			right = p[1]
		}

	}
	return res
}
