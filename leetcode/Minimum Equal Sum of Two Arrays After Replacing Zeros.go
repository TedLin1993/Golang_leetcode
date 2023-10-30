package leetcode

func minSum(nums1 []int, nums2 []int) int64 {
	min1, zero1 := 0, 0
	for _, num := range nums1 {
		if num > 0 {
			min1 += num
		} else {
			min1 += 1
			zero1++
		}
	}
	min2, zero2 := 0, 0
	for _, num := range nums2 {
		if num > 0 {
			min2 += num
		} else {
			min2 += 1
			zero2++
		}
	}
	if zero1 == 0 && min1 < min2 {
		return -1
	}
	if zero2 == 0 && min2 < min1 {
		return -1
	}
	return int64(max(min1, min2))
}
