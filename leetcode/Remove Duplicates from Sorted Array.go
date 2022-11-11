package leetcode

func removeDuplicates(nums []int) int {
	left := 0
	for right := 1; right < len(nums); right++ {
		if nums[right] != nums[left] {
			left++
			nums[left] = nums[right]
		}
	}
	return left + 1
}
