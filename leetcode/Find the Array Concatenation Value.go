package leetcode

import "fmt"

func findTheArrayConcVal(nums []int) int64 {
	res := 0
	left, right := 0, len(nums)-1
	for left < right {
		lefeValue, rightValue := nums[left], nums[right]
		for rightValue != 0 {
			lefeValue = lefeValue * 10
			rightValue = rightValue / 10
		}
		res += lefeValue + nums[right]
		left++
		right--
	}
	if left == right {
		res += nums[left]
	}
	return int64(res)
}

func Test_findTheArrayConcVal() {
	fmt.Println(findTheArrayConcVal([]int{7, 52, 2, 4}))      //596
	fmt.Println(findTheArrayConcVal([]int{5, 14, 13, 8, 12})) //673
}
