package leetcode

func minOrAfterOperations(nums []int, k int) int {
	res := 0
	mask := 0
	for i := 29; i >= 0; i-- {
		mask |= 1 << i
		cnt := 0
		cur := mask
		for _, v := range nums {
			cur &= mask & v
			if cur != 0 {
				cnt++
			} else {
				cur = mask
			}
		}
		if cnt > k {
			mask ^= 1 << i
			res |= 1 << i
		}
	}
	return res
}
