package leetcode

import "fmt"

func trap(heights []int) int {
	res := 0
	left, right := 0, len(heights)-1
	leftMax, rightMax := heights[left], heights[right]
	for left <= right {
		if leftMax < rightMax {
			if heights[left] > leftMax {
				leftMax = heights[left]
			} else {
				res += leftMax - heights[left]
			}
			left++
		} else {
			if heights[right] > rightMax {
				rightMax = heights[right]
			} else {
				res += rightMax - heights[right]
			}
			right--
		}
	}
	return res
}

func Test_trap() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})) //6
	fmt.Println(trap([]int{4, 2, 0, 3, 2, 5}))                   //9
}
