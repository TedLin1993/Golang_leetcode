package leetcode

func maximumStrongPairXor(nums []int) int {
	n := len(nums)
	res := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if abs(nums[i]-nums[j]) <= min(nums[i], nums[j]) {
				res = max(res, nums[i]^nums[j])
			}
		}
	}
	return res
}
