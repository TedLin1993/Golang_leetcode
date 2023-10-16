package leetcode

func findIndices_2905(nums []int, indexDifference int, valueDifference int) []int {
	n := len(nums)
	minIdx, maxIdx := 0, 0
	for i := indexDifference; i < n; i++ {
		if nums[i-indexDifference] < nums[minIdx] {
			minIdx = i - indexDifference
		}
		if nums[i-indexDifference] > nums[maxIdx] {
			maxIdx = i - indexDifference
		}
		if nums[i]-nums[minIdx] >= valueDifference {
			return []int{minIdx, i}
		}
		if nums[maxIdx]-nums[i] >= valueDifference {
			return []int{maxIdx, i}
		}
	}
	return []int{-1, -1}
}
