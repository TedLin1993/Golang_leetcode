package leetcode

import (
	"fmt"
	"math"
)

func distributeCookies(cookies []int, k int) int {
	n := len(cookies)
	children := make([]int, k)
	res := math.MaxInt
	var backTracking func(idx int, currentMax int)
	backTracking = func(idx int, currentMax int) {
		if currentMax >= res {
			return
		}
		if idx == n {
			res = currentMax
			return
		}

		for i := range children {
			children[i] += cookies[idx]
			backTracking(idx+1, max(currentMax, children[i]))
			children[i] -= cookies[idx]
		}
	}
	backTracking(0, 0)
	return res
}

func Test_distributeCookies() {
	fmt.Println(distributeCookies([]int{8, 15, 10, 20, 8}, 2)) //31
}
