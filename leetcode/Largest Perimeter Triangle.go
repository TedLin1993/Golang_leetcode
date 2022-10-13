package leetcode

import (
	"fmt"
	"sort"
)

func largestPerimeter(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	for i := 0; i < len(nums)-2; i++ {
		if nums[i] < nums[i+1]+nums[i+2] {
			return nums[i] + nums[i+1] + nums[i+2]
		}
	}
	return 0
}

func Test_largestPerimeter() {
	fmt.Println(largestPerimeter([]int{2, 1, 2}))    //5
	fmt.Println(largestPerimeter([]int{1, 2, 1}))    //0
	fmt.Println(largestPerimeter([]int{3, 6, 2, 3})) //8
}
