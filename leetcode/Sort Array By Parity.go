package leetcode

import "fmt"

func sortArrayByParity(nums []int) []int {
	firstOddIndex := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			nums[firstOddIndex], nums[i] = nums[i], nums[firstOddIndex]
			firstOddIndex++
		}
	}
	return nums
}

func TestsortArrayByParity() {
	fmt.Println(sortArrayByParity([]int{3, 1, 2, 4})) //[2 4 3 1]
	fmt.Println(sortArrayByParity([]int{0}))          //[0]
}
