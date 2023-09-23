package leetcode

import "sort"

func longestStrChain(words []string) int {
	n := len(words)
	sort.Slice(words, func(w1, w2 int) bool {
		return len(words[w1]) < len(words[w2])
	})
	dp := make([]int, n)
	res := 1
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := i - 1; j >= 0; j-- {
			if dp[j]+1 > dp[i] && isPredecessor(words[j], words[i]) {
				dp[i] = dp[j] + 1
			}
		}
		res = max(res, dp[i])
	}
	return res
}

func isPredecessor(w1, w2 string) bool {
	if len(w1)+1 != len(w2) {
		return false
	}
	w1Idx := 0
	for i := 0; i < len(w2); i++ {
		if w1[w1Idx] == w2[i] {
			w1Idx++
		}
		if w1Idx == len(w1) {
			return true
		}
	}
	return w1Idx == len(w1)
}
