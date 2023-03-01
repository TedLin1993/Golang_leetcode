package leetcode

func leftRigthDifference(nums []int) []int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	res := make([]int, len(nums))
	left := 0
	right := sum
	for i := 0; i < len(nums); i++ {
		right -= nums[i]
		value := left - right
		if value < 0 {
			value = -value
		}
		res[i] = value
		left += nums[i]
	}
	return res
}
