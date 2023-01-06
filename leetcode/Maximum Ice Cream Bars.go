package leetcode

import (
	"fmt"
	"sort"
)

func maxIceCream(costs []int, coins int) int {
	sort.Ints(costs)
	res := 0
	for i := 0; i < len(costs); i++ {
		if costs[i] <= coins {
			res++
			coins -= costs[i]
		} else {
			break
		}
	}
	return res
}

func Test_maxIceCream() {
	fmt.Println(maxIceCream([]int{1, 3, 2, 4, 1}, 7))     //4
	fmt.Println(maxIceCream([]int{10, 6, 8, 7, 7, 8}, 5)) //0
	fmt.Println(maxIceCream([]int{1, 6, 3, 1, 2, 5}, 20)) //6
}
