package leetcode

import "sort"

func maximumStrongPairXor_2935(nums []int) int {
	sort.Ints(nums)
	mask := 0
	res := 0
	for i := 19; i >= 0; i-- {
		mask |= 1 << i
		nextRes := res | 1<<i
		dict := map[int]int{}
		for _, y := range nums {
			mask_y := y & mask
			if x, ok := dict[nextRes^mask_y]; ok && 2*x >= y {
				res = nextRes
				break
			}
			dict[mask_y] = y
		}
	}
	return res
}
