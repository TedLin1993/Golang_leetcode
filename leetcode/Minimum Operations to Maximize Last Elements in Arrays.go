package leetcode

func minOperations_2934(nums1 []int, nums2 []int) int {
	n := len(nums1)
	countOp := func(last1, last2 int) int {
		res := 0
		for i := 0; i < n-1; i++ {
			if nums1[i] > last1 || nums2[i] > last2 {
				if nums1[i] > last2 || nums2[i] > last1 {
					return n + 1
				}
				res++
			}
		}
		return res
	}
	res := min(countOp(nums1[n-1], nums2[n-1]), 1+countOp(nums2[n-1], nums1[n-1]))
	if res > n {
		return -1
	}
	return res
}
