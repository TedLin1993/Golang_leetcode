package leetcode

func longestEqualSubarray(nums []int, k int) int {
	n := len(nums)
	numIdx := make([][]int, n+1)
	for i, num := range nums {
		numIdx[num] = append(numIdx[num], i)
	}
	res := 1
	for _, idxs := range numIdx {
		if len(idxs) == 0 {
			continue
		}
		delCount := 0
		left, right := 0, 0
		for right < len(idxs) {
			if right > left {
				delCount += idxs[right] - idxs[right-1] - 1
				for delCount > k {
					delCount -= idxs[left+1] - idxs[left] - 1
					left++
				}
			}
			res = max(res, right-left+1)
			right++
		}
	}
	return res
}
