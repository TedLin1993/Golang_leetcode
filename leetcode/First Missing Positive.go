package leetcode

import (
	"fmt"
)

func firstMissingPositive(nums []int) int {
	index := 0
	//We can store position in-place, so space complexity is O(1)
	//Let nums length is n, we replace nums to [1,2,3,...,n]
	//Then we iterate the new nums, we can easily find the missing positive value
	for index < len(nums) {
		//out of bound, ignore
		if nums[index] <= 0 || nums[index] > len(nums) {
			index++
			continue
		}

		correctIdx := nums[index] - 1
		//already in corrected position, ignore
		if nums[index] == index+1 || nums[correctIdx] == correctIdx+1 {
			index++
			continue
		}

		nums[index], nums[correctIdx] = nums[correctIdx], nums[index]
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1
}

func Test_firstMissingPositive() {
	fmt.Println(firstMissingPositive([]int{1, 2, 0}))         //3
	fmt.Println(firstMissingPositive([]int{3, 4, -1, 1}))     //2
	fmt.Println(firstMissingPositive([]int{7, 8, 9, 11, 12})) //1
	fmt.Println(firstMissingPositive([]int{1, 2, 3}))         //4
	fmt.Println(firstMissingPositive([]int{1, 1}))            //2
}
