package leetcode

import (
	"fmt"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLen := len(nums1) + len(nums2)
	isEven := totalLen%2 == 0
	idx1, idx2 := 0, 0
	res, prev := 0, 0
	for i := 0; i <= totalLen/2; i++ {
		prev = res
		if idx1 == len(nums1) {
			res = nums2[idx2]
			idx2++
		} else if idx2 == len(nums2) {
			res = nums1[idx1]
			idx1++
		} else if nums1[idx1] < nums2[idx2] {
			res = nums1[idx1]
			idx1++
		} else {
			res = nums2[idx2]
			idx2++
		}
	}

	if isEven {
		return (float64)(res+prev) / 2
	}
	return (float64)(res)
}

func TestMedianOfTwoSortedArrays() {
	var nums1 []int
	var nums2 []int

	nums1 = []int{1, 3}
	nums2 = []int{2}
	fmt.Println(findMedianSortedArrays(nums1, nums2)) //2

	nums1 = []int{0, 0}
	nums2 = []int{0, 0}
	fmt.Println(findMedianSortedArrays(nums1, nums2)) //0

	nums1 = []int{1, 2}
	nums2 = []int{3, 4}
	fmt.Println(findMedianSortedArrays(nums1, nums2)) //2.5

	nums1 = []int{1}
	nums2 = []int{}
	fmt.Println(findMedianSortedArrays(nums1, nums2)) //1
}
