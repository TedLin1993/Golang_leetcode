package leetcode

func findIndices(nums []int, indexDifference int, valueDifference int) []int {
	n := len(nums)
	for i := range nums {
		for j := i + indexDifference; j < n; j++ {
			if abs(nums[i]-nums[j]) >= valueDifference {
				return []int{i, j}
			}
		}
	}
	return []int{-1, -1}
}
