package leetcode

func maximumSumOfHeights(maxHeights []int) int64 {
	n := len(maxHeights)
	res := 0
	for i := 0; i < n; i++ {
		cur := maxHeights[i]
		prev := maxHeights[i]
		for j := i - 1; j >= 0; j-- {
			v := min(maxHeights[j], prev)
			cur += v
			prev = v
		}
		prev = maxHeights[i]
		for j := i + 1; j < n; j++ {
			v := min(maxHeights[j], prev)
			cur += v
			prev = v
		}
		res = max(res, cur)
	}
	return int64(res)
}
