package leetcode

func singleNonDuplicate(nums []int) int {
	length := len(nums)
	left, right := 0, length-1
	for left < right {
		mid := left + (right-left)/2
		value := nums[mid]
		if mid == 0 || mid == length-1 {
			return value
		}
		if nums[mid-1] != value && nums[mid+1] != value {
			return value
		}
		if nums[mid-1] == value {
			if mid%2 != 0 {
				left = mid + 1
			} else {
				right = mid - 2
			}
		} else {
			if (mid+1)%2 != 0 {
				left = mid + 2
			} else {
				right = mid - 1
			}
		}
	}
	return nums[left]
}
