package leetcode

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	if len(nums) == 3 {
		return max(nums[1], nums[0]+nums[2])
	}
	dp0, dp1, dp2 := nums[0], nums[1], nums[0]+nums[2]
	for i := 3; i < len(nums); i++ {
		temp := max(dp0, dp1) + nums[i]
		dp0, dp1, dp2 = dp1, dp2, temp
	}
	return max(dp1, dp2)
}
