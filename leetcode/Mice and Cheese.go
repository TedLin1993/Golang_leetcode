package leetcode

import (
	"fmt"
	"sort"
)

func miceAndCheese(reward1 []int, reward2 []int, k int) int {
	sum, n := 0, len(reward1)
	diff := make([]int, n)
	for i := 0; i < n; i++ {
		sum += reward2[i]
		diff[i] = reward1[i] - reward2[i]
	}
	sort.Ints(diff)
	for i := n - 1; i > n-1-k; i-- {
		sum += diff[i]
	}
	return sum
}

func Test_miceAndCheese() {
	fmt.Println(miceAndCheese([]int{1, 1, 3, 4}, []int{4, 4, 1, 1}, 2))       //15
	fmt.Println(miceAndCheese([]int{1, 1}, []int{1, 1}, 2))                   //2
	fmt.Println(miceAndCheese([]int{4, 1, 5, 3, 3}, []int{3, 4, 4, 5, 2}, 3)) //21
	fmt.Println(miceAndCheese([]int{10, 10}, []int{1, 1}, 2))                 //20
	fmt.Println(miceAndCheese([]int{1}, []int{4}, 1))                         //1
}
