package leetcode

import (
	"fmt"
	"math"
)

func find132pattern(nums []int) bool {
	stack := []int{}
	thirdElement := math.MinInt32
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < thirdElement {
			return true
		}
		if len(stack) == 0 {
			stack = append(stack, nums[i])
			continue
		}
		if nums[i] > stack[len(stack)-1] {
			for len(stack) > 0 && nums[i] > stack[len(stack)-1] {
				thirdElement = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, nums[i])
		} else if nums[i] < stack[len(stack)-1] {
			stack = append(stack, nums[i])
		}
	}

	return false
}

func Testfind132pattern() {
	fmt.Println(find132pattern([]int{1, 2, 3, 4}))                                                                         // false
	fmt.Println(find132pattern([]int{3, 1, 4, 2}))                                                                         // true
	fmt.Println(find132pattern([]int{-1, 3, 2, 0}))                                                                        // true
	fmt.Println(find132pattern([]int{3, 5, 0, 3, 4}))                                                                      // true
	fmt.Println(find132pattern([]int{4, 2, 3, 1}))                                                                         // false
	fmt.Println(find132pattern([]int{3, 4, 1, 1, 1, 1, 1, 1, 2}))                                                          // false
	fmt.Println(find132pattern([]int{0, -1000, 2000, -3000, 4000, -5000, 6000, -7000, 8000, -9000, 10000, -11000, 12000})) // false
}
