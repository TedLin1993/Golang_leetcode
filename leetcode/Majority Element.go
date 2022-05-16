package leetcode

import "fmt"

func majorityElement(nums []int) int {
	candidate := nums[0]
	count := 1
	for _, num := range nums[1:] {
		if num == candidate {
			count += 1
		} else {
			if count > 0 {
				count -= 1
			} else {
				count = 1
				candidate = num
			}
		}
	}
	return candidate
}

func TestMajorityElement() {
	fmt.Println(majorityElement([]int{3, 2, 3}))             //3
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2})) //2
}
