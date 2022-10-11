package leetcode

import (
	"fmt"
	"math"
)

func increasingTriplet(nums []int) bool {
	first, second := math.MaxInt64, math.MaxInt64
	for i := 0; i < len(nums); i++ {
		if nums[i] > second {
			return true
		}
		if nums[i] < first {
			first = nums[i]
		} else if nums[i] > first && nums[i] < second {
			second = nums[i]
		}
	}

	return false
}

func Test_increasingTriplet() {
	fmt.Println(increasingTriplet([]int{1, 2, 3, 4, 5}))    //true
	fmt.Println(increasingTriplet([]int{5, 4, 3, 2, 1}))    //false
	fmt.Println(increasingTriplet([]int{2, 1, 5, 0, 4, 6})) //true
}
