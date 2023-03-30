package leetcode

import (
	"fmt"
	"sort"
)

func maxSatisfaction(satisfaction []int) int {
	sort.Ints(satisfaction)
	res, sum := 0, 0
	for i := len(satisfaction) - 1; i >= 0; i-- {
		if sum+satisfaction[i] > 0 {
			sum += satisfaction[i]
			res += sum
		} else {
			break
		}
	}
	return res
}

func Test_maxSatisfaction() {
	fmt.Println(maxSatisfaction([]int{-1, -8, 0, 5, -9})) //14
	fmt.Println(maxSatisfaction([]int{4, 3, 2}))          //20
	fmt.Println(maxSatisfaction([]int{-1, -4, -5}))       //0
}
