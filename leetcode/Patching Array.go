package leetcode

func minPatches(nums []int, n int) int {
	s := 1
	i := 0
	res := 0
	for s <= n {
		if i < len(nums) && nums[i] <= s {
			s += nums[i]
			i++
		} else {
			s *= 2
			res++
		}
	}
	return res
}
