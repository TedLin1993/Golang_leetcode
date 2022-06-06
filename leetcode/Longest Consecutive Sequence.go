package leetcode

import "fmt"

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	numMap := make(map[int]bool, len(nums))
	for _, v := range nums {
		numMap[v] = true
	}

	res := 1
	for _, v := range nums {
		if numMap[v-1] {
			continue
		}

		temp := 1
		high := v + 1
		for numMap[high] {
			temp++
			high++
		}

		if temp > res {
			res = temp
		}
	}
	return res
}

func TestlongestConsecutive() {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))         //4
	fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1})) //9
}
