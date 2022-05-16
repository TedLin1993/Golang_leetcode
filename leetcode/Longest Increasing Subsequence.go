package leetcode

import "fmt"

func lengthOfLIS(nums []int) int {
	var res []int
	res = append(res, nums[0])

	for i := 1; i < len(nums); i++ {
		temp := nums[i]
		if temp > res[len(res)-1] {
			res = append(res, temp)
		} else {
			j := binarySearch(res, temp)
			res[j] = temp
		}
	}
	fmt.Println("res=", res)
	return len(res)
}

func binarySearch(a []int, temp int) int {
	left := 0
	right := len(a) - 1
	var mid int

	for left < right {
		mid = (left + right) / 2
		if a[mid] == temp {
			return mid
		}

		if a[mid] < temp {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}

func TestLengthOfLIS() {
	fmt.Println(lengthOfLIS([]int{0, 3, 1, 6, 2, 2, 5})) //4
	// fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18})) //4
	// fmt.Println(lengthOfLIS([]int{0, 1, 0, 3, 2, 3}))           //4
	// fmt.Println(lengthOfLIS([]int{7, 7, 7, 7, 7, 7, 7}))        //1
	// fmt.Println(lengthOfLIS([]int{5, 9, 2, 10})) //4

}
