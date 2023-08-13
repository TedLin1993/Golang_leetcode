package leetcode

import "fmt"

func validPartition(nums []int) bool {
	n := len(nums)
	isValid := func(nums []int) bool {
		if len(nums) == 2 {
			return nums[0] == nums[1]
		}
		if len(nums) == 3 {
			return (nums[0] == nums[1] && nums[1] == nums[2]) || (nums[0]+1 == nums[1] && nums[0]+2 == nums[2])
		}
		return false
	}

	if n <= 3 {
		return isValid(nums)
	}

	dp1, dp2 := isValid(nums[:2]), isValid(nums[:3])
	for i := 3; i < n; i++ {
		dp := (dp1 && isValid(nums[i-2:i+1])) || (dp2 && isValid(nums[i-1:i+1]))
		dp1, dp2 = dp2, dp
	}
	fmt.Println(dp1, dp2)
	return dp2
}

func Test_validPartition() {
	fmt.Println(validPartition([]int{4, 4, 4, 5, 6})) //true
}
