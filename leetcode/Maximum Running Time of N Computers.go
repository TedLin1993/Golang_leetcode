package leetcode

import "fmt"

func maxRunTime(n int, batteries []int) int64 {
	right := 0
	for _, b := range batteries {
		right += b
	}
	right = right / n
	left := 1
	for left < right {
		target := left + (right-left+1)/2
		sum := 0
		for _, b := range batteries {
			sum += min(b, target)
		}
		if sum >= n*target {
			left = target
		} else {
			right = target - 1
		}
	}
	return int64(left)
}

func Test_maxRunTime() {
	fmt.Println(maxRunTime(12, []int{11, 89, 16, 32, 70, 67, 35, 35, 31, 24, 41, 29, 6, 53, 78, 83})) //43
}
