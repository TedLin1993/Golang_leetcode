package leetcode

import (
	"fmt"
	"sort"
)

func maxIncreasingGroups(usageLimits []int) int {
	n := len(usageLimits)
	sort.Slice(usageLimits, func(i, j int) bool {
		return usageLimits[i] > usageLimits[j]
	})
	max := n
	for i := 0; i < n; i++ {
		if usageLimits[i] < n-i {
			max = min(max, usageLimits[i]+i)
		}
	}
	return max
}

func Test_maxIncreasingGroups() {
	fmt.Println(maxIncreasingGroups([]int{1, 2, 5})) //3
	fmt.Println(maxIncreasingGroups([]int{1, 2, 2})) //2
}
