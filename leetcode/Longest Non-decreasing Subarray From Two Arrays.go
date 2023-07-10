package leetcode

import "fmt"

func maxNonDecreasingLength(nums1 []int, nums2 []int) int {
	n := len(nums1)
	res := 1
	dp1, dp2 := 1, 1
	for i := 1; i < n; i++ {
		temp1, temp2 := dp1, dp2
		dp1, dp2 = 1, 1
		if nums1[i] >= nums1[i-1] {
			dp1 = temp1 + 1
		}
		if nums1[i] >= nums2[i-1] {
			dp1 = max(dp1, temp2+1)
		}
		res = max(res, dp1)

		if nums2[i] >= nums1[i-1] {
			dp2 = temp1 + 1
		}
		if nums2[i] >= nums2[i-1] {
			dp2 = max(dp2, temp2+1)
		}
		res = max(res, dp2)
	}

	return res
}

func Test_maxNonDecreasingLength() {
	fmt.Println(maxNonDecreasingLength([]int{2, 3, 1}, []int{1, 2, 1}))                   //2
	fmt.Println(maxNonDecreasingLength([]int{1, 3, 2, 1}, []int{2, 2, 3, 4}))             //4
	fmt.Println(maxNonDecreasingLength([]int{8, 7, 4}, []int{13, 4, 4}))                  //2
	fmt.Println(maxNonDecreasingLength([]int{11, 7, 7, 9}, []int{19, 19, 1, 7}))          //3
	fmt.Println(maxNonDecreasingLength([]int{4, 2}, []int{10, 4}))                        //2
	fmt.Println(maxNonDecreasingLength([]int{8, 15, 14, 13, 19}, []int{18, 8, 5, 4, 10})) //3
}
