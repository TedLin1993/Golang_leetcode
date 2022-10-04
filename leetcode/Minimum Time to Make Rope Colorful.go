package leetcode

import "fmt"

func minCost(colors string, neededTime []int) int {
	res := 0
	prevColor := byte(' ')
	max := 0
	for i := 0; i < len(colors); i++ {
		if colors[i] != prevColor {
			max = neededTime[i]
			prevColor = colors[i]
		} else {
			if neededTime[i] > max {
				res += max
				max = neededTime[i]
			} else {
				res += neededTime[i]
			}
		}
	}
	return res
}

func Test_minCost() {
	fmt.Println(minCost("abaac", []int{1, 2, 3, 4, 5})) //3
	fmt.Println(minCost("abc", []int{1, 2, 3}))         //0
	fmt.Println(minCost("aabaa", []int{1, 2, 3, 4, 1})) //2
}
