package leetcode

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	numIndex := m + n - 1
	for m > 0 && n > 0 {
		if nums1[m-1] > nums2[n-1] {
			nums1[numIndex] = nums1[m-1]
			m--
		} else {
			nums1[numIndex] = nums2[n-1]
			n--
		}
		numIndex--
	}
	for n > 0 {
		nums1[numIndex] = nums2[n-1]
		numIndex--
		n--
	}
}

func Testmerge() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	merge(nums1, 3, []int{2, 5, 6}, 3)
	fmt.Println(nums1)
}
