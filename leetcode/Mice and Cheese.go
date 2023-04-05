package leetcode

import (
	"fmt"
)

func miceAndCheese(reward1 []int, reward2 []int, k int) int {
	res := 0

	n := len(reward1)
	visited := make([]bool, n)
	for k > 0 {
		amount := 0
		idx1, idx2 := 0, 0
		for i := 0; i < n; i++ {
			if visited[i] {
				continue
			}
			for j := 0; j < n; j++ {
				if visited[j] || j == i {
					continue
				}
				if reward1[i]+reward2[j] > amount {
					amount = reward1[i] + reward2[j]
					idx1, idx2 = i, j
				}
			}
		}
		res += amount
		visited[idx1] = true
		visited[idx2] = true
		k--
	}
	return res
}

func Test_miceAndCheese() {
	// fmt.Println(miceAndCheese([]int{1, 1, 3, 4}, []int{4, 4, 1, 1}, 2))       //15
	// fmt.Println(miceAndCheese([]int{1, 1}, []int{1, 1}, 2))                   //2
	fmt.Println(miceAndCheese([]int{4, 1, 5, 3, 3}, []int{3, 4, 4, 5, 2}, 3)) //21
	fmt.Println(miceAndCheese([]int{10, 10}, []int{1, 1}, 2))                 //20
}
